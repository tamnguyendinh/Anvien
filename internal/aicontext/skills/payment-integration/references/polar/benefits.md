# Polar Benefits

Current guidance for Polar automated benefits, benefit grants, license keys, credits, feature flags, file downloads, GitHub/Discord access, custom benefits, and entitlement sync.

## Source Links

- Benefits introduction: https://polar.sh/docs/features/benefits/introduction
- Credits: https://polar.sh/docs/features/usage-based-billing/credits
- Feature Flags: https://polar.sh/docs/features/benefits/feature-flags
- License Keys: https://polar.sh/docs/features/benefits/license-keys
- Customer State: https://polar.sh/docs/integrate/customer-state

## Benefit Types

| Benefit | Use For | Key Integration Point |
|---------|---------|-----------------------|
| Credits | Prepaid usage units for a meter | Usage billing and Customer State meter balance |
| License Keys | Software/API access keys | Customer Portal license key endpoints |
| Feature Flags | Simple app-side entitlement flags | Customer State or `customer.state_changed` |
| File Downloads | Secure digital downloads | Customer Portal |
| GitHub Access | Private repository access | Polar-managed GitHub invitation/revocation |
| Discord Access | Discord invite/roles | Polar-managed Discord invite/role handling |
| Custom | Manual or app-specific fulfillment | App logic plus benefit grant events |

Benefits are standalone resources attached to products. A single benefit can be attached to multiple products.

## Benefit Grants

Benefit grants represent a customer receiving a benefit. They are created, updated, and revoked as purchases, renewals, refunds, cancellations, and product changes happen.

Webhook events:

- `benefit_grant.created`
- `benefit_grant.updated`
- `benefit_grant.revoked`
- `customer.state_changed`

Use `customer.state_changed` when you want one entitlement sync event that includes active subscriptions, granted benefits, and meter balances.

## Credits Benefit

Credits are the official way to prepay usage units for a meter.

Behavior:

- Subscription products credit the customer's meter balance at the beginning of each subscription cycle.
- One-time products credit the customer once at purchase.
- Credits are consumed before overage charges.
- If you want credits-only spending without overage invoices, do not attach a metered price for that meter.
- Polar does not block usage when credits run out; your app must enforce usage gates.

Load `usage-based-billing.md` for event ingestion and customer meter balance checks.

## Feature Flag Benefit

Feature Flag benefits are lightweight entitlement flags. If the customer has an active grant for the benefit, the feature is available.

Use cases:

- Premium feature gates.
- Beta access.
- Tier-specific limits.
- API quota tiers.

Recommended access check:

1. Fetch Customer State or handle `customer.state_changed`.
2. Find the active benefit grant for the feature flag benefit ID.
3. Apply app-side access or quota based on the benefit metadata.

## License Keys

License key benefits support:

- Brandable prefixes.
- Automatic expiration.
- Activation limits.
- Custom validation conditions.
- Usage quota increments.
- Automatic revocation when subscriptions are canceled or benefits are revoked.

Validation guidance:

- Validate using organization ID.
- If activation limits are enabled, activate before validation and pass activation ID.
- If multiple license-key benefits exist in one organization, validate the returned `benefit_id` so a key for one product cannot unlock another product.
- If using usage limits, increment usage during validation only for real billable usage.

## File Downloads

Use file download benefits for downloadable digital files. Customers access downloads through the Customer Portal. Keep fulfillment instructions clear and include versioning notes in the product/customer communication.

## GitHub Access

GitHub access benefits can invite customers to private repositories and revoke access automatically when entitlement ends.

Implementation notes:

- Keep required permissions minimal.
- Use separate benefits for distinct access tiers.
- Still listen to grants/customer state if your app mirrors repository access locally.

## Discord Access

Discord benefits can invite customers and grant roles.

Implementation notes:

- Ensure the Polar Discord app is configured.
- Design role hierarchy before attaching products.
- Do not duplicate role revocation logic if Polar is authoritative.

## Custom Benefits

Use Custom benefits for manual or app-specific fulfillment that Polar does not automate.

Examples:

- Priority support.
- Manual onboarding.
- Private community not supported by a built-in integration.
- A Cal.com link or service credit handled by your own backend.

Custom benefit grants still emit grant events and appear in Customer State.

## Product Attachment

Benefits can be attached from the dashboard or API. When benefits change:

- Existing eligible customers may receive newly attached benefits.
- Removed benefits can be revoked.
- Webhooks/customer state should be used to sync local entitlement state.

## Entitlement Strategy

For most SaaS apps:

1. Attach benefits to products in Polar.
2. Set `external_customer_id` on checkout/customer creation.
3. Listen to `customer.state_changed`.
4. Store the latest Customer State snapshot or normalized entitlement rows.
5. Gate features from local state for request-time speed.
6. Periodically reconcile critical access from Polar.

Avoid inferring entitlements only from subscription status when benefits are attached. Customer State is usually more complete.
