export const asBigInt = (n: bigint | number | string) => {
  if (typeof n === "bigint") {
    return n;
  }
  return BigInt(n);
};
