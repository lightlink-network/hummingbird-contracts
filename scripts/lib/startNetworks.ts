import { exec, ChildProcess, execSync } from "child_process";

interface StartNetworkOptions {
  logOutput?: boolean;
}

const killProcessOnPort = (port: number): void => {
  try {
    const processIds = execSync(`lsof -t -i:${port}`)
      .toString()
      .trim()
      .split("\n");
    processIds.forEach((processId) => {
      if (processId) {
        execSync(`kill -9 ${processId}`);
      }
    });
  } catch (error) {
    // Ignore errors if no process is running on the port
  }
};

const startNetwork = (
  port: number,
  networkName: string,
  options?: StartNetworkOptions,
): Promise<ChildProcess> => {
  return new Promise((resolve, reject) => {
    console.log(`    Starting ${networkName} network on port ${port}...`);
    const process = exec(
      `npx hardhat node --port ${port}`,
      (error, stdout, stderr) => {
        if (error) {
          console.error(
            `    Error starting ${networkName} network: ${error.message}`,
          );
          reject(error);
        }
      },
    );

    if (options?.logOutput) {
      // Capture and log the output from the process
      process.stdout?.on("data", (data) => {
        console.log(`[${networkName}] ${data.toString()}`);
      });

      process.stderr?.on("data", (data) => {
        console.error(`[${networkName} ERROR] ${data.toString()}`);
      });
    }

    resolve(process);
  });
};

const retryStartNetwork = async (
  port: number,
  networkName: string,
  options?: StartNetworkOptions,
  retries: number = 5,
): Promise<ChildProcess> => {
  for (let attempt = 1; attempt <= retries; attempt++) {
    try {
      return await startNetwork(port, networkName, options);
    } catch (error) {
      console.log(
        `    Retrying to start ${networkName} network on port ${port}... (Attempt ${attempt} of ${retries})`,
      );
      killProcessOnPort(port);
      await new Promise((resolve) => setTimeout(resolve, 1000)); // Wait for a second before retrying
    }
  }
  throw new Error(
    `    Failed to start ${networkName} network on port ${port} after ${retries} attempts.`,
  );
};

export const startNetworks = async (
  options?: StartNetworkOptions,
): Promise<{ l1Network: ChildProcess; l2Network: ChildProcess }> => {
  killProcessOnPort(8545);
  killProcessOnPort(8546);

  console.log("    Starting networks...");

  try {
    const l1Network = await retryStartNetwork(8545, "l1", options);
    console.log("    L1 network started successfully on port 8545.");

    const l2Network = await retryStartNetwork(8546, "l2", options);
    console.log("    L2 network started successfully on port 8546.");

    // Allow some time for Hardhat networks to start
    console.log("    Waiting for networks to be fully operational...");
    await new Promise((resolve) => setTimeout(resolve, 5000));

    console.log("    Networks are up and running.\n\n");
    return { l1Network, l2Network };
  } catch (error) {
    console.error("    Error starting networks: ", error);
    throw error;
  }
};
