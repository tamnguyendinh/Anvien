# Desktop App Reference

> Đọc file này khi dự án là desktop application (Electron, Tauri, hoặc native).
> Bổ sung vào các bước tương ứng trong SKILL.md — không thay thế chúng.

---

## Chọn Framework Desktop

| Framework | Stack | App size | Performance | Phù hợp khi |
|-----------|-------|----------|-------------|-------------|
| **Electron** | Node.js + Chromium | ~150MB | Trung bình | Team web, cần ecosystem npm đầy đủ |
| **Tauri** | Rust + WebView hệ thống | ~5MB | Cao | Muốn bundle nhỏ, security tốt hơn |
| **Neutralino** | C++ + WebView | ~2MB | Cao | App đơn giản, không cần Node backend |
| **Flutter Desktop** | Dart | ~20MB | Cao | Cross-platform (mobile + desktop cùng codebase) |
| **native (Swift/WinUI/Qt)** | Platform native | Nhỏ nhất | Tốt nhất | Platform-specific, team có kinh nghiệm native |

> **Rule of thumb:** SaaS desktop companion app → Tauri. Internal tool / phức tạp → Electron. Mobile-first mở rộng desktop → Flutter.

---

## Bước 2 — Architecture (bổ sung cho Desktop)

```markdown
### Desktop Architecture Patterns

#### Main Process / Core Process (Electron: main, Tauri: Rust core)
- Quản lý cửa sổ (BrowserWindow / WebviewWindow)
- Truy cập filesystem, OS APIs, native features
- IPC handler — nhận lệnh từ renderer/frontend
- System tray, notifications, auto-launch, deep link

#### Renderer Process / Frontend (WebView)
- UI layer — React/Vue/Svelte như web bình thường
- KHÔNG gọi trực tiếp Node/OS APIs
- Giao tiếp qua IPC bridge duy nhất

#### IPC Contract (critical — phải document đầy đủ)
| Channel / Command | Direction | Payload | Response |
|-------------------|-----------|---------|----------|
| `app:get-version` | R → M | none | string |
| `file:open-dialog`| R → M | filter[] | string[] paths |
| `settings:save`   | R → M | Settings | void |
| `update:available`| M → R | UpdateInfo | void |

Direction: R = Renderer→Main, M = Main→Renderer (push event)

#### Local Data Strategy
- User settings: OS config dir (app.getPath('userData'))
- App data / cache: app.getPath('appData')
- Documents: user-chosen path, nhớ trong settings
- KHÔNG hardcode paths — luôn dùng app.getPath() hoặc Tauri path API

#### Online/Offline Mode
- App có hoạt động offline không? Đến mức nào?
- Sync strategy khi reconnect: last-write-wins / conflict resolution
- Offline queue: lưu actions vào local DB, flush khi online
```

---

## Bước 3 — Library & Toolchain (bổ sung cho Desktop)

```markdown
### Desktop Framework
- Framework: Electron v32 / Tauri v2 / ...
- Bundler: electron-builder / electron-forge / Tauri CLI
- UI: React 19 + Vite (trong WebView)

### Electron-specific
- IPC: contextBridge + ipcRenderer (KHÔNG dùng remote module)
- Preload script: expose typed API từ main sang renderer
- Security: contextIsolation: true, nodeIntegration: false (bắt buộc)
- Auto-update: electron-updater (Squirrel) / Tauri updater

### Tauri-specific
- Commands: #[tauri::command] trong Rust, invoke() từ frontend
- Permissions: tauri.conf.json — khai báo explicit (fs, shell, http, ...)
- Plugins: tauri-plugin-store (settings), tauri-plugin-sql (local DB), ...

### Local Storage / Database
- Nhẹ: electron-store / @tauri-apps/plugin-store (key-value JSON)
- Nặng: better-sqlite3 (Electron) / tauri-plugin-sql + SQLite (Tauri)
- Secrets / keychain: keytar (Electron) / tauri-plugin-stronghold (Tauri)

### OS Integration
- System tray: Electron Tray API / tray plugin Tauri
- Notifications: node-notifier / Tauri notification plugin
- Deep links: protocol registration (myapp://...)
- Auto-launch: auto-launch npm / Tauri autostart plugin
- Global shortcuts: globalShortcut (Electron) / global-shortcut plugin

### Testing Desktop
- Unit: Vitest (frontend) + Jest/Vitest (main process logic)
- E2E Desktop: Playwright với Electron driver / WebdriverIO + Tauri WebDriver
```

---

## Bước 4 — Local Database (SQLite pattern)

```markdown
### SQLite cho Desktop

#### Khi nào dùng SQLite thay vì JSON files
- Dữ liệu có quan hệ (relational)
- Cần query, filter, sort phức tạp
- Dataset > vài nghìn records
- Cần transactions (atomic operations)

#### Schema considerations cho desktop
- Thêm cột `synced_at TIMESTAMP` nếu có cloud sync
- Thêm cột `local_id TEXT` (ULID/UUID) — không phụ thuộc server-assigned ID
- `is_dirty BOOLEAN DEFAULT 0` — đánh dấu record chưa sync

#### Migration local DB
- Lưu schema version trong user's DB
- Migration chạy tự động khi app start
- KHÔNG phá vỡ schema cũ — chỉ ADD columns, không DROP/RENAME
- Backup DB trước khi migrate (copy file .db sang .db.bak)
```

---

## Bước 7 — DevOps & Distribution (Desktop-specific)

```markdown
### Build & Sign

#### macOS
- Code signing: Developer ID Application certificate (Apple Developer)
- Notarization: bắt buộc từ macOS 10.15+ (xcrun notarytool)
- Targets: x64 (Intel), arm64 (Apple Silicon), universal binary
- Format: .dmg (distribution) + .app (app bundle)

#### Windows
- Code signing: EV certificate (Extended Validation) để tránh SmartScreen warning
- Targets: x64, arm64 (tuỳ chọn)
- Format: .exe installer (NSIS / WiX) + .msi (enterprise)

#### Linux
- Format: .deb (Debian/Ubuntu), .rpm (Fedora), .AppImage (universal), .snap
- No signing required nhưng nên có GPG signature cho repo

### Auto-Update Strategy
- Update server: GitHub Releases / S3 / custom server
- Update channel: stable / beta / nightly
- Update flow:
  1. App start → check for update (silent)
  2. Update available → notify user (không force)
  3. User confirm → download in background
  4. Downloaded → prompt restart
  5. Restart → apply update
- Rollback: giữ lại version trước, cho phép rollback nếu lỗi
- Delta updates: electron-differentiator / Tauri NSIS nightly (giảm bandwidth)

### Release Pipeline
```
tag vX.Y.Z →
  CI: lint + test + build (mac/win/linux parallel) →
  Sign & Notarize (macOS) →
  Sign (Windows) →
  Upload to GitHub Releases / S3 →
  Publish update manifest →
  Smoke test (download + install + launch)
```

### Distribution Channels
| Channel | Effort | Reach | Notes |
|---------|--------|-------|-------|
| Direct download (website) | Thấp | Trung bình | Cần signing để tránh warning |
| Mac App Store | Cao | Cao | Sandbox restrictions, review time |
| Microsoft Store | Trung bình | Trung bình | MSIX packaging |
| Homebrew Cask | Thấp | Dev-focused | Community-driven |
| Winget | Thấp | Dev-focused | Microsoft package manager |

### App Updates & Versioning
- SemVer: MAJOR.MINOR.PATCH
- Build number: CI build number (monotonically increasing)
- macOS: CFBundleVersion phải tăng dần, CFBundleShortVersionString = semver
- Windows: FileVersion phải tăng dần
```

---

## Security Checklist (Desktop)

```markdown
### Electron Security
- [ ] contextIsolation: true
- [ ] nodeIntegration: false
- [ ] sandbox: true (Electron 20+)
- [ ] webSecurity: true (không disable)
- [ ] allowRunningInsecureContent: false
- [ ] Preload script: chỉ expose những gì cần thiết qua contextBridge
- [ ] Validate mọi IPC message từ renderer (không trust renderer)
- [ ] Không eval() content từ user / remote
- [ ] CSP header cho BrowserWindow

### Tauri Security
- [ ] Khai báo explicit permissions trong tauri.conf.json
- [ ] Scope filesystem access (không cho phép truy cập toàn bộ FS)
- [ ] Validate input trước khi pass vào Rust commands
- [ ] Không dùng shell execute với user input

### Credentials & Secrets
- KHÔNG lưu secret trong plain text / localStorage / electron-store không encrypted
- Dùng OS keychain: keytar (Electron), tauri-plugin-stronghold (Tauri)
- API keys: gọi qua backend server, không hardcode trong app bundle
- Token storage: encrypted với OS keychain, clear on logout

### Update Security
- Verify update signature trước khi apply
- HTTPS only cho update server
- Certificate pinning nếu cần
```

---

## Đặc thù Test cho Desktop

```markdown
### Test Levels

#### Unit (main process / Rust commands)
- Test business logic trong isolation
- Mock IPC, filesystem, OS APIs
- Tools: Vitest (Node), cargo test (Rust)

#### Integration (IPC layer)
- Test main process handlers với fake renderer calls
- Verify IPC contract khớp với frontend expectations
- Mock OS dialogs, notifications

#### E2E Desktop
- Electron: Playwright + @playwright/test với electronApp
  ```typescript
  const electronApp = await electron.launch({ args: ['main.js'] });
  const window = await electronApp.firstWindow();
  await window.click('#login-btn');
  ```
- Tauri: WebdriverIO + tauri-driver
- Test flows: install, first launch, settings save/restore, update flow

#### Manual Smoke Test Checklist (pre-release)
- [ ] Install từ fresh (không có previous version)
- [ ] Upgrade từ previous version (data migration)
- [ ] Uninstall (không để lại rác trong system)
- [ ] Launch lần đầu (onboarding)
- [ ] Offline mode hoạt động đúng
- [ ] Auto-update flow (staging channel)
- [ ] Deep link handling
- [ ] System tray hoạt động
```