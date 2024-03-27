// log to stderr
export const log = (...msg: any[]) => {
  process.stderr.write(msg.map((m) => JSON.stringify(m)).join(" ") + "\n");
};
