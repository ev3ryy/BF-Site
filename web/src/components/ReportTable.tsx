import React, { useMemo, useState } from 'react'
import {
  useReactTable,
  ColumnDef,
  ColumnFiltersState,
  getCoreRowModel,
  getFilteredRowModel,
  getSortedRowModel,
  flexRender,
} from '@tanstack/react-table'

export interface Report {
  id: number
  adminName: string
  role: string
  punishDate: string
  description: string
  evidence: string
}

interface ReportsTableProps {
  data: Report[]
}

const ReportsTable: React.FC<ReportsTableProps> = ({ data }) => {
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([])

  const columns = useMemo<ColumnDef<Report, any>[]>(
    () => [
      { accessorKey: 'adminName', header: 'Admin Name', filterFn: 'includesString' },
      { accessorKey: 'role',      header: 'Role',       filterFn: 'includesString' },
      { accessorKey: 'punishDate', header: 'Punish Date', filterFn: 'includesString' },
      { accessorKey: 'description', header: 'Description', filterFn: 'includesString' },
      { accessorKey: 'evidence',   header: 'Evidence',    filterFn: 'includesString' },
    ],
    []
  )

  const table = useReactTable({
    data,
    columns,
    state: { columnFilters },
    onColumnFiltersChange: setColumnFilters,
    getCoreRowModel: getCoreRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    getSortedRowModel: getSortedRowModel(),
  })

  return (
    <div className="overflow-x-auto shadow-md rounded-lg">
      {/* Фильтры */}
      <div className="p-4 grid grid-cols-2 gap-4">
        {table.getAllLeafColumns().map(column => (
          <input
            key={column.id}
            type="text"
            value={(column.getFilterValue() ?? '') as string}
            onChange={e => column.setFilterValue(e.target.value)}
            placeholder={`Filter ${column.columnDef.header}`}
            className="border rounded p-2 w-full"
          />
        ))}
      </div>

      <table className="min-w-full bg-white table-auto">
        <thead className="bg-gray-100">
          {table.getHeaderGroups().map(headerGroup => (
            <tr key={headerGroup.id}>
              {headerGroup.headers.map(header => (
                <th
                  key={header.id}
                  colSpan={header.colSpan}
                  className={`p-3 text-left select-none ${header.column.getCanSort() ? 'cursor-pointer' : ''}`}
                  onClick={header.column.getToggleSortingHandler()}
                >
                  {flexRender(header.column.columnDef.header, header.getContext())}
                  <span className="ml-2">
                    {{ asc: '🔼', desc: '🔽' }[header.column.getIsSorted() as string] ?? null}
                  </span>
                </th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody>
          {table.getRowModel().rows.map(row => (
            <tr key={row.id} className="hover:bg-gray-50">
              {row.getVisibleCells().map(cell => (
                <td key={cell.id} className="border px-4 py-2">
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

export default ReportsTable