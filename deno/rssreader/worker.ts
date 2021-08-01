import { parseFeed, bgBlue, bold } from "./deps.ts";

self.onmessage = async (e) => {
  const { url } = e.data;

  const response = await fetch(url);
  const xml = await response.text();
  const feed = await parseFeed(xml);
  let counter = 0;

  console.log("-------------------");
  console.log(bgBlue(bold(url)));
  for (const entry of feed.entries) {
    if (counter === 5) {
      break
    }

    if (entry.title !== undefined) {
      console.log(entry.title?.value)
      counter++
    }
  }

  self.close();
};

