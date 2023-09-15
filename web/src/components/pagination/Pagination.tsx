import React from 'react';
import {
    Button,
    Flex, NumberDecrementStepper,
    NumberIncrementStepper,
    NumberInput,
    NumberInputField,
    NumberInputStepper, Select
} from "@chakra-ui/react";
import { VscChevronLeft, VscChevronRight } from "react-icons/vsc";
import {useBoundStore} from "../../stores";
const SimplePagination = () => {
    const store = useBoundStore((state) => state)
    const {page, limit, totalPage, total} = store.pagination
    const {setPage, setLimit} = store
    return(
        <Flex gap={2} >
            <Button variant={'outline'}>
                <VscChevronLeft />
            </Button>
            <NumberInput width={100} defaultValue={page+1} min={1} max={20}>
                <NumberInputField />
            </NumberInput>
            <Button variant={'outline'}>
                <VscChevronRight />
            </Button>
            <Select
                defaultValue={limit}
                onChange={(e) => {
                 let value = e.target.value
                    if(value === 'all'){
                        setLimit('')
                    } else {
                        setLimit(value)
                    }
                }}
                placeholder='Select option' width={100}>
                <option value={10}>10</option>
                <option value='50'>50</option>
                <option value='100'>100</option>
                <option value='1000'>1000</option>
                <option value='all'>All</option>
            </Select>
        </Flex>
    )
}
export default SimplePagination;