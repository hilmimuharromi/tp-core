import {StateCreator} from "zustand";
import axios from "axios";
import {UserStore} from "./user";
// import pagination from "../components/pagination/Pagination";
import {GlobalStore} from "./global";

export interface SuppliersStore {
    productSuppliers: []
    pagination: PaginationType
    setPage: (page: number) => void
    setLimit: (limit: string) => void
    getProductsSuppliers: (spCode: string) => void
    filterProductSupplier: FilterProductSupplier
    setFilterProductSupplier: (value: string, key: string)  => void
}

type PaginationType = {
    page: number
    limit: number
    totalPage: number
    total: number
}

type FilterProductSupplier = {
    name: string
    operator: string
    category: string
}


export const createSuppliersStore:StateCreator<SuppliersStore & UserStore & GlobalStore, [], [], SuppliersStore> = ((setState, getState, ) => (
    {
        productSuppliers: [],
        pagination: {
            page: 0,
            limit: 10,
            totalPage: 1,
            total:10
        },
        filterProductSupplier: {
            name: "",
            category: "",
            operator:""
        },
        setPage: (page) => setState(() => ({pagination: {
            ...getState().pagination, page: page
            }})),
        setLimit: (limit) => setState(() => ({pagination: {
            ...getState().pagination, limit : Number(limit)
            }
        })),
        setFilterProductSupplier: (value, key) => setState(() => ({filterProductSupplier: {
                ...getState().filterProductSupplier, [key] : value
            }
        })),
        getProductsSuppliers: async () => {
            getState().setLoading(true)
            const token = getState().user.token
            const page = getState().pagination.page
            const limit = getState().pagination.limit
            const {name, operator, category} = getState().filterProductSupplier
            const {data} = await axios.get(`/api/v1/suppliers/products?page=${page}&limit=${limit}&name=${name}&category=${category}&operator=${operator}`,{
                headers: {
                    Authorization: `Bearer ${token}`
                }
            })
            if(data.code == 200) {
                setState(() => (
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