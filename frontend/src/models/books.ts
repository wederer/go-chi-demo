const mode = import.meta.env.MODE

export interface Model {
  id: string
  title: string,
  no_pages: number,
}

/**
 * Calls mockedFn if the app is running in E2E mode. Otherwise calls fn.
 *
 * @param mockedFn
 * @param fn
 */
function mockIfE2E<T>(mockedFn: () => Promise<T>, fn: () => Promise<T>): Promise<T> {
  // mode is passed to vite cli as --mode e2e
  // so we can mock the call to the backend
  if (mode === 'e2e') {
    return mockedFn()
  }

  return fn()
}


export async function deleteOne(id: string): Promise<void> {
  return mockIfE2E<void>(deleteOneMocked, _deleteOne(id))
}

async function deleteOneMocked() {
  // to simulate errors in the backend
  // in "real" E2E tests we would use the actual backend
  if (window.__E2E_ERROR__) {
    throw new Error()
  }
  return
}

function _deleteOne(id: string) {
  return async function () {
    const response = await fetch(`http://localhost:3000/books/${id}`, {method: 'DELETE'})
    if (!response.ok) {
      throw new Error(response.statusText);
    }
  }
}

export async function getAll(): Promise<Array<Model>> {
  return mockIfE2E<Array<Model>>(getAllMocked, _getAll)
}

async function getAllMocked() {
  // to simulate errors in the backend
  // in "real" E2E tests we would use the actual backend
  if (window.__E2E_ERROR__) {
    throw new Error()
  }
  return [{id: "some-id", title: "some-title", no_pages: 42}, {id: "some-other-id", title: "some-other-title", no_pages: 43}]
}

async function _getAll() {
  const response = await fetch("http://localhost:3000/books")
  if (!response.ok) {
    throw new Error(response.statusText);
  }

  return await response.json()
}
