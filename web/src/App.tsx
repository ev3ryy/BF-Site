import React, { useState, useEffect } from 'react'
import ReportsTable, { Report } from './components/ReportTable'

const App: React.FC = () => {
  const [reports, setReports] = useState<Report[]>([])
  const [loading, setLoading] = useState<boolean>(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const fetchReports = async () => {
      try {
        const response = await fetch('http://localhost:8080/reports')
        if (!response.ok) throw new Error(`HTTP ${response.status}`)
        const data: Array<{
          id: number
          adminName: string
          roleName: string
          punishDate: string
          description: string
          evidence: string
        }> = await response.json()
        const mapped: Report[] = data.map((r) => ({
          id: r.id,
          adminName: r.adminName,
          role: r.roleName,
          punishDate: r.punishDate,
          description: r.description,
          evidence: r.evidence,
        }))
        setReports(mapped)
      } catch (err: any) {
        console.error(err)
        setError(err.message)
      } finally {
        setLoading(false)
      }
    }
    fetchReports()
  }, [])

  if (loading) return <div className="text-center p-4">Загрузка...</div>
  if (error) return <div className="text-center p-4 text-red-500">Ошибка: {error}</div>

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Система отчётов</h1>
      <ReportsTable data={reports} />
    </div>
  )
}

export default App
