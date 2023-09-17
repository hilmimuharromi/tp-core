import React from 'react';
import {Flex, Select} from "@chakra-ui/react";
import {SearchInput} from 'components/input'
import {useBoundStore} from "../../../../stores";
const FilterTableSupplierProduct = () => {
    const store = useBoundStore((state) => state)
    const {setFilterProductSupplier, filterProductSupplier} = store
    const {name,category, operator} = filterProductSupplier
    return(
        <Flex gap={2}>
            <SearchInput
                value={name}
                onChange={(e: any) => setFilterProductSupplier(e.target.value, "name")}/>
            <Select
                width={'200px'}
                placeholder='operator'
                value={operator}
                onChange={(e: any) => setFilterProductSupplier(e.target.value, "operator")}
            >
                <option value={'telkomsel'}>Telkomsel</option>
                <option value='indosat'>Indosat</option>
                <option value='three'>Three</option>
                <option value='xl'>XL</option>
                <option value='axis'>Axis</option>
            </Select>
            <Select
                width={'200px'}
                placeholder='category'
            value={category}
                onChange={(e: any) => setFilterProductSupplier(e.target.value, "category")}

            >
                <option value={'data'}>Data</option>
                <option value='pulsa'>Pulsa</option>
                <option value='pulsa'>Pulsa Transfer</option>
                <option value='pula'>Pulsa Voucher</option>
                <option value='pln'>PLN</option>
                <option value='game'>Game</option>
                <option value='e-wallet'>e-Wallet</option>
                <option value='e-money'>e-Money</option>
            </Select>
        </Flex>
    )
}
export default FilterTableSupplierProduct;