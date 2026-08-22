export function run(value: Promise<number>): Promise<number> {
  const maximum = Math.max(value ? 1 : 0, 2)
  return new Promise<number>((resolve) => resolve(maximum))
}
