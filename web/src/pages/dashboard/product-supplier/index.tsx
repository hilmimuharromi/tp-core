import {Box, Flex} from '@chakra-ui/react';
import ProductsTable from 'pages/dashboard/product-supplier/components/ProductsSuppliersTable';
// import CheckTable from 'views/admin/dataTables/components/CheckTable';
// import ColumnsTable from 'views/admin/dataTables/components/ColumnsTable';
// import ComplexTable from 'views/admin/dataTables/components/ComplexTable';
// import tableDataDevelopment from 'views/admin/dataTables/variables/tableDataDevelopment';
// import tableDataCheck from 'views/admin/dataTables/variables/tableDataCheck';
// import tableDataColumns from 'views/admin/dataTables/variables/tableDataColumns';
// import tableDataComplex from 'views/admin/dataTables/variables/tableDataComplex';
import {useBoundStore} from "../../../stores";
import {useEffect} from "react";

export default function ProductSupplier() {
    const store = useBoundStore((state) => state)
    const {getProductsSuppliers, productSuppliers, loading} = store
    const {limit, page} = store.pagination
    const {name,operator, category} = store.filterProductSupplier

    useEffect(() => {
        getProductsSuppliers('ds')
        console.log('ps ===>', productSuppliers)
    }, [page, limit, name, operator, category])

    if (loading) return <Flex>Loading .....</Flex>

    return (
        <Box pt={{ base: '130px', md: '80px', xl: '80px' }}>
            <ProductsTable   tableData={productSuppliers ? productSuppliers : []} />
        </Box>
    );
}
