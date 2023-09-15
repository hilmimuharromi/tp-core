import {create} from 'zustand'
import {createSuppliersStore, SuppliersStore} from './suppliers'
import {createUserStore, UserStore} from "./user";
import {createGlobalStore, GlobalStore} from "./global";
import { devtools, persist } from 'zustand/middleware'
export const useBoundStore =
    create<GlobalStore & UserStore & SuppliersStore>()(devtools(
        persist(
            (...params) => ({
                ...createSuppliersStore(...params),
                ...createUserStore(...params),
                ...createGlobalStore(...params)
            }),
            {
                name: 'tp-core-storage'
            }
        )
    ))