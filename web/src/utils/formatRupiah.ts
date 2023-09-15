export const formatRupiah = (value: number) => {
    if(value) {
        return new Intl.NumberFormat("id-ID", { style: "currency", currency: "IDR",minimumFractionDigits: 0 }).format(value)
    } else return new Intl.NumberFormat("id-ID", { style: "currency", currency: "IDR",minimumFractionDigits: 0 }).format(0)

}