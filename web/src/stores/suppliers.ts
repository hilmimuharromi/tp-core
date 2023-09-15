import {StateCreator} from "zustand";
import axios from "axios";
import {UserStore} from "./user";
import pagination from "../components/pagination/Pagination";
import {GlobalStore} from "./global";

export interface SuppliersStore {
    productSuppliers: []
    pagination: PaginationType
    setPage: (page: number) => void
    setLimit: (limit: string) => void
    getProductsSuppliers: (spCode: string) => void
}

type PaginationType = {
    page: number
    limit: number
    totalPage: number
    total: number
}


export const createSuppliersStore:StateCreator<SuppliersStore & UserStore & GlobalStore, [], [], SuppliersStore> = ((setState, getState, store) => (
    {
        productSuppliers: [],
        pagination: {
            page: 0,
            limit: 10,
            totalPage: 1,
            total:10
        },
        setPage: (page) => setState(() => ({pagination: {
            ...getState().pagination, page: page
            }})),
        setLimit: (limit) => setState(() => ({pagination: {
            ...getState().pagination, limit : Number(limit)
            }
        })),
        getProductsSuppliers: async (spCode) => {
            getState().setLoading(true)
            const token = getState().user.token
            const page = getState().pagination.page
            const limit = getState().pagination.limit
            const {data} = await axios.get(`/api/v1/suppliers/products?page=${page}&limit=${limit}`,{
                headers: {
                    Authorization: `Bearer ${token}`
                }
            })
            if(data.code == 200) {
                setState((state) => (
                    {
                        productSuppliers: data.data,
                        pagination: {
                            limit: limit,
                            page: page,
                            totalPage: data.meta.totalPage,
                            total: data.meta.total
                        }
                    }))
            }
            getState().setLoading(false)
            console.log("data get products", data.data)
        },
    }
))