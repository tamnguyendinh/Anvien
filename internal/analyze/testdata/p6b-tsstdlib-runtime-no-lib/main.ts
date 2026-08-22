export function run(value: Promise<number>): Promise<number> {
  return Math.max(value ? 1 : 0, 2)
}
