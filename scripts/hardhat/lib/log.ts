// log to stderr
export const log = (...msg: any[]) => {
  process.stderr.write(
    msg
      .map((m) => {
        if (typeof m === "string") return m;
        if (m.toString) return m.toString();
        if (typeof m === "object") return JSON.stringify(m, null, 2);
        return JSON.stringify(m);
      })
      .join(" ") + "\n",
  );
};
