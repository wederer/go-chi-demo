declare global {
  interface Window  {
    // used to simulate backend errors
    __E2E_ERROR__: boolean
  }
}

// so that TS recognizes this file as a module this is needed
// https://www.typescriptlang.org/docs/handbook/release-notes/typescript-1-8.html#example-6
export {}