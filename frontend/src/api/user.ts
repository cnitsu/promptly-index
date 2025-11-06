export const login = async (username: string, password: string) => {
  if (username === 'guess' && password === 'guess') {
    return { code: 0, data: { prompt: "{:username}@promptly $", username: "guess" } }
  }
  return { code: 1, msg: "username or password error", data: { prompt: "{:username}@promptly $", username: "guess" } }
}
export const getLoginUser = async () => {
  return { code: 10, data: { prompt: "test", username: "fuck" } }// TODO
}
