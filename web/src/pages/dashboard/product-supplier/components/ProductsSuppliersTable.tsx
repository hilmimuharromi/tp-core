import {
    Box,
    Flex,
    Progress,
    Table, TableCaption,
    Tbody,
    Td,
    Text,
    Th,
    Thead,
    Tr,
    useColorModeValue
} from '@chakra-ui/react';
import {
    createColumnHelper,
    flexRender,
    getCoreRowModel,
    getSortedRowModel,
    SortingState,
    useReactTable,
} from '@tanstack/react-table';
// Custom components
import Card from 'components/card/Card';
// import Menu from 'components/menu/MainMenu';
import {formatRupiah} from 'utils/formatRupiah'
import SimplePagination from 'components/pagination/Pagination'
import FilterTable from './FilterTableSupplierProduct'
// import { AndroidLogo, AppleLogo, WindowsLogo } from 'components/icons/Icons';
import * as React from 'react';
// Assets

type RowObj = {
    name: string;
    operator: string;
    status: string;
    price: number;
    category: string
};

const columnHelper = createColumnHelper<RowObj>();

// const columns = columnsDataCheck;
export default function ComplexTable(props: { tableData: any, isFetching: boolean }) {
    const { tableData, isFetching=false } = props;
    const [ sorting, setSorting ] = React.useState<SortingState>([]);
    const textColor = useColorModeValue('secondaryGray.900', 'white');
    // const iconColor = useColorModeValue('secondaryGray.500', 'white');
    const borderColor = useColorModeValue('gray.200', 'whiteAlpha.100');
    let defaultData = tableData;
    const columns = [
        columnHelper.accessor('name', {
            id: 'name',
            header: () => (
                <Text
                    justifyContent='space-between'
                    align='center'
                    fontSize={{ sm: '10px', lg: '12px' }}
                    color='gray.400'>
                    NAME
                </Text>
            ),
            cell: (info: any) => (
                <Flex align='center'>
                    <Text color={textColor} fontSize='sm' fontWeight='700'>
                        {info.getValue()}
                    </Text>
                </Flex>
            )
        }),
        columnHelper.accessor('operator', {
            id: 'operator',
            header: () => (
                <Text
                    justifyContent='space-between'
                    align='center'
                    fontSize={{ sm: '10px', lg: '12px' }}
                    color='gray.400'>
                    Operator
                </Text>
            ),
            cell: (info: any) => (
                <Flex align='center'>
                    <Text color={textColor} fontSize='sm' fontWeight='300'>
                        {info.getValue()}
                    </Text>
                </Flex>
            )
        }),
        columnHelper.accessor('category', {
            id: 'category',
            header: () => (
                <Text
                    justifyContent='space-between'
                    align='center'
                    fontSize={{ sm: '10px', lg: '12px' }}
                    color='gray.400'>
                    Category
                </Text>
            ),
            cell: (info: any) => (
                <Flex align='center'>
                    <Text color={textColor} fontSize='sm' fontWeight='300'>
                        {info.getValue()}
                    </Text>
                </Flex>
            )
        }),

        columnHelper.accessor('status', {
            id: 'status',
            header: () => (
                <Text
                    justifyContent='space-between'
                    align='center'
                    fontSize={{ sm: '10px', lg: '12px' }}
                    color='gray.400'>
                    Status
                </Text>
            ),
            cell: (info) => (
                <Flex align='center'>
                    <Text me='10px' color={textColor} fontSize='sm' fontWeight='700'>
                        {info.getValue()}
                    </Text>
                </Flex>
            )
        }),
        columnHelper.accessor('price', {
            id: 'price',
            header: () => (
                <Text
                    justifyContent='space-between'
                    align='center'
                    fontSize={{ sm: '10px', lg: '12px' }}
                    color='gray.400'>
                    Price
                </Text>
            ),
            cell: (info) => (
                <Flex align='center'>
                    <Text me='10px' color={textColor} fontSize='sm' fontWeight='700'>
                        {formatRupiah(info.getValue())}
                    </Text>
                </Flex>
            )
        })
    ];
    const [ data ] = React.useState(() => [ ...defaultData ]);
    const table = useReactTable({
        data:  tableData && tableData.length > 0 ? tableData : [],
        columns,
        state: {
            sorting
        },
        autoResetAll: isFetching,
        onSortingChange: setSorting,
        getCoreRowModel: getCoreRowModel(),
        getSortedRowModel: getSortedRowModel(),
        debugTable: true,
        manualPagination: true
    });

    // if (isFetching) return <Stack>
    //     <Skeleton height='20px' />
    //     <Skeleton height='20px' />
    //     <Skeleton height='20px' />
    // </Stack>
    return (
        <Card flexDirection='column' w='100%' px='0px' overflowX={{ sm: 'scroll', lg: 'hidden' }}>
            <Flex px='25px' mb="8px" justifyContent='space-between' align='center'>
                <Text color={textColor} fontSize='22px' fontWeight='700' lineHeight='100%'>
                    Supplier Product
                </Text>
                <FilterTable />
                {/*<Menu />*/}
            </Flex>
            <Box>
                <Table variant='simple' color='gray.500' mb='24px' mt="12px">
                    <Thead>
                        {table.getHeaderGroups().map((headerGroup) => (
                            <Tr key={headerGroup.id}>
                                {headerGroup.headers.map((header) => {
                                    return (
                                        <Th
                                            key={header.id}
                                            colSpan={header.colSpan}
                                            pe='10px'
                                            borderColor={borderColor}
                                            cursor='pointer'
                                            onClick={header.column.getToggleSortingHandler()}>
                                            <Flex
                                                justifyContent='space-between'
                                                align='center'
                                                fontSize={{ sm: '10px', lg: '12px' }}
                                                color='gray.400'>
                                                {flexRender(header.column.columnDef.header, header.getContext())}{{
                                                asc: '',
                                                desc: '',
                                            }[header.column.getIsSorted() as string] ?? null}
                                            </Flex>
                                        </Th>
                                    );
                                })}
                            </Tr>
                        ))}
                    </Thead>
                    {
                        isFetching ? (
                            <TableCaption>
                                <Progress size='sm' isIndeterminate />
                            </TableCaption>
                        ) : !!tableData.length ? (
                            <Tbody>
                                {
                                    !isFetching && table.getRowModel().rows.slice(0).map(row => {
                                        return (
                                            <Tr key={row.id}>
                                                {row.getVisibleCells().map((cell) => {
                                                    return (
                                                        <Td
                                                            key={cell.id}
                                                            fontSize={{ sm: '14px' }}
                                                            minW={{ sm: '150px', md: '200px', lg: 'auto' }}
                                                            borderColor='transparent'>
                                                            {flexRender(cell.column.columnDef.cell, cell.getContext())}
                                                        </Td>
                                                    );
                                                })}
                                            </Tr>
                                        )
                                    })
                                }
                            </Tbody>
                        ) : (
                            <TableCaption>
                                <Text>There are no data to display.</Text>
                            </TableCaption>
                        )
                    }

                </Table>
                <Flex  mt={10} width={"100%"} justifyContent={'end'}>
                    <SimplePagination />
                </Flex>
            </Box>
        </Card>
    );
}
