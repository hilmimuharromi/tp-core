import {StateCreator} from "zustand";
import axios from "axios";
import {GlobalStore} from "./global";

export interface UserStore {
    user: UserType,
    login: (email: string, password: string) => void
}

type UserType = {
    id: string
    name: string
    email: string
    token: string
}
export const createUserStore:StateCreator<UserStore & GlobalStore, [], [], UserStore> = ((set, get) => (
    {
        user: {
            id: "",
            name: "",
            email: "",
            token: "",
        },
        login: async (email, password) => {
            get().setLoading(true)
            const {data} = await axios.post("/api/v1/user/login", {
                email: email,
                password: password
            })
            // setTimeout(() => {
            //     get().setLoading(false)
            // }, 2000)
            set((state) => ({
                user: data.data,
                loading: false,
                isLogin: true
            }))
            console.log("data login", data)
        },
    }
))