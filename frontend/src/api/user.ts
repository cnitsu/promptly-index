import instance from "../utils/axios"

export const login = async (username: string, password: string) => {
  if (username === 'guess' && password === 'guess') {
    return { code: 0, data: { prompt: "{:username}@promptly $", username: "guess" } }
  }
  return await instance.post('/v1/user/login', { username: username, password: password })
    .then((response) => {
      console.log(response.data);
      return response.data.data
    })
}
export const getLoginUser = async () => {
  return await instance.get('/v1/user/info')
    .then((response) => {
      return response.data.data
    })
}
