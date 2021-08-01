if (Deno.args.length < 1) {
  console.error("please specify filename");
  Deno.exit(1);
}

const data = await Deno.readTextFile(Deno.args[0]);
for (const url of data.split("\n")) {
  if (url.length !== 0 ) {
    const worker = new Worker(new URL("./worker.ts", import.meta.url).href, { type: "module" });
    worker.postMessage({ url: url });
  }
}
