import { defineStore } from "pinia";
import { getLoginUser } from "../api/user";

export const useUserStore = defineStore('user', {
  state: () => ({
    loginUser: {
    },
  }),
  getters: {},
  // persist: {
  //   key: "user-store",
  //   storage: window.localStorage,
  //   beforeRestore: (ctx) => {
  //     console.log("load userStore data start")
  //   },
  //   afterRestore: (ctx) => {
  //     console.log("load userStore data end");
  //   },
  // },
  actions: {
    async getAndSetLoginUser() {
      const res = await getLoginUser()
      if (res?.code === 0 && res.data) {
        this.loginUser = res.data
      } else {
        this.$reset()
      }
    },
    setLoginUser(user: promptly.UserType) {
      this.loginUser = user
    },
  },
})
