import {defineStore} from 'pinia'

import {getToken, getUser, removeToken, removeUser, setToken, setUser} from "@/utils/auth.js";
import request from "@/utils/request.js";


function getDefaultState() {
    return {
        token: getToken(),
        name: getUser(),
    }
}

export const useUserStore = defineStore('user', {
    state: () => {
        return getDefaultState()
    },
    actions: {
        resetState() {
            this.token = getToken();
            this.name = getUser();
        },
        setToken(token) {
            this.token = token;
        },
        setName(name) {
            this.name = name;
        },
        async login(userInfo) {
            const {username, password} = userInfo;
            const resp = await request.post("/api/user/login", {username: username, password: password})
            const token = resp.data
            this.setToken(token)
            this.setName(username)

            setToken(token)
            setUser(username)

        },
        async logout() {
            try{
                const res = await request.get("/api/user/logout", this.token);
                removeToken()
                removeUser()
                this.resetState();
            } catch(err){
                console.log(err)
            }

        },
        resetToken() {
            removeToken();
            this.resetState();
        }
    },
})
