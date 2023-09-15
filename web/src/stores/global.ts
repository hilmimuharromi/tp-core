import {StateCreator} from "zustand";
import axios from "axios";
export interface GlobalStore {
    loading: boolean,
    isLogin: boolean,
    setLoading: (status: boolean) => void
}

export const createGlobalStore:StateCreator<GlobalStore> = ((set) => (
    {
        loading: false,
        isLogin: false,
        setLoading: (status) => set((state) => ({loading: status}))
    }
))