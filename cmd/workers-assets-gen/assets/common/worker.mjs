import mod from "./app.wasm";
import * as imports from "./shim.mjs";

imports.init(mod);
const oldFetch = global.fetch;
global.fetch = (input, init) => {
  if (init) {
    init.credentials = undefined;
  }
  return oldFetch(input, init);
};
const oldFetch2 = globalThis.fetch;
globalThis.fetch = (input, init) => {
  if (init) {
    init.credentials = undefined;
  }
  return oldFetch2(input, init);
};
export default { fetch: imports.fetch, scheduled: imports.scheduled };
