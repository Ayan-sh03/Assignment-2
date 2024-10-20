// src/pages/ChartsPage.js
import { ChartData, ChartOptions } from 'chart.js'
import 'chart.js/auto'
import { useEffect, useState } from 'react'
import { Bar, Line, Pie } from 'react-chartjs-2'

interface WeatherData {
  date: string

  city: string
  maxTemp: number
  minTemp: number
  avgTemp: number
  dominantCondition: string
}

interface AlertData {
  city: string
  // date: string
  alertCount: number
}
function ChartsPage(): JSX.Element {
  const [weatherData, setWeatherData] = useState<WeatherData[]>([])
  const [alertsData, setAlertsData] = useState<AlertData[]>([])
  const [selectedCity, setSelectedCity] = useState<string>()
  const [cities, setCities] = useState<string[]>([])
  useEffect(() => {
    //get cities from local storage
    const cities = JSON.parse(localStorage.getItem('cities') || '[]')
    setCities(cities)
    setSelectedCity(cities[0])
    fetchWeatherData(selectedCity?.toLowerCase())
    fetchAlertsData()

    //eslint-disable-next-line react-hooks/exhaustive-deps
  }, [selectedCity])

  const fetchWeatherData = async (city: string | undefined) => {
    // Mock weather data for the selected city
    if (city == '') city = 'nagpur'

    const res = await fetch(
      `http://localhost:8080/summary/${city}/${
        new Date().toISOString().split('T')[0]
      }`
    )
    const data = await res.json()

    const weatherData: WeatherData[] = data.map((item: any) => ({
      date: new Date().toISOString().split('T')[0],
      maxTemp: item.max_temperature,
      minTemp: item.min_temperature,
      avgTemp: item.average_temperature,
      dominantCondition: item.dominant_condition,
      city: selectedCity,
    }))

    setWeatherData(weatherData)
  }

  const fetchAlertsData = async () => {
    const res = await fetch(`http://localhost:8080/alerts`)
    const data = await res.json()
    console.log('printing ', data)

    const alertData = data.map((item: any) => ({
      city: item.city.charAt(0).toUpperCase() + item.city.slice(1),
      alertCount: item.count,
    }))

    setAlertsData(alertData)
  }

  // Prepare data for charts
  const dates = weatherData.map((data) => data.date)
  const maxTemps = weatherData.map((data) => data.maxTemp)
  const minTemps = weatherData.map((data) => data.minTemp)
  const avgTemps = weatherData.map((data) => data.avgTemp)
  const dominantConditions = weatherData.reduce<Record<string, number>>(
    (acc, data) => {
      acc[data.dominantCondition] = (acc[data.dominantCondition] || 0) + 1
      return acc
    },
    {}
  )

  const alertsPerCity = alertsData.reduce<Record<string, number>>(
    (acc, data) => {
      acc[data.city] = (acc[data.city] || 0) + data.alertCount
      return acc
    },
    {}
  )

  const lineChartData: ChartData<'line'> = {
    labels: dates,
    datasets: [
      {
        label: 'Max Temperature (°C)',
        data: maxTemps,
        borderColor: 'rgba(255, 99, 132, 1)',
        backgroundColor: 'rgba(255, 99, 132, 0.2)',
      },
    ],
  }

  const barChartData: ChartData<'bar'> = {
    labels: dates,
    datasets: [
      {
        label: 'Average Temperature (°C)',
        data: avgTemps,
        backgroundColor: 'rgba(75, 192, 192, 0.6)',
      },
    ],
  }

  const pieChartData: ChartData<'pie'> = {
    labels: Object.keys(dominantConditions),
    datasets: [
      {
        data: Object.values(dominantConditions),
        backgroundColor: [
          '#FF6384',
          '#36A2EB',
          '#FFCE56',
          '#66BB6A',
          '#BA68C8',
        ],
      },
    ],
  }

  const alertsBarChartData: ChartData<'bar'> = {
    labels: Object.keys(alertsPerCity),
    datasets: [
      {
        label: 'Alert Count',
        data: Object.values(alertsPerCity),

        backgroundColor: 'rgba(255, 159, 64, 0.6)',
      },
    ],
  }

  const alertsBarChartOptions: ChartOptions<'bar'> = {
    indexAxis: 'y',
    scales: {
      x: {
        beginAtZero: true,
        suggestedMax: Math.max(...Object.values(alertsPerCity)) + 3, // Add some buffer
      },
    },
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-50 p-8">
      <div className="max-w-7xl mx-auto">
        {/* Header Section */}
        <div className="flex justify-between items-center mb-8">
          <div>
            <h1 className="text-3xl font-bold text-gray-800">
              Weather Analytics Dashboard
            </h1>
            <p className="text-gray-600 mt-2">
              Real-time weather insights for Indian cities
            </p>
          </div>

          <select
            value={selectedCity}
            onChange={(e) => setSelectedCity(e.target.value)}
            className="px-4 py-2 bg-white border border-gray-200 rounded-lg shadow-sm
                     focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500
                     text-gray-700 hover:border-gray-300 transition-colors duration-200 capitalize"
          >
            {cities.map((city) => (
              <option key={city} value={city} className="capitalize">
                {city}
              </option>
            ))}
          </select>
        </div>

        {/* Quick Stats */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
          {/* Max Temperature Card */}
          <div className="bg-white/70 backdrop-blur-sm rounded-xl shadow-lg hover:shadow-xl transition-shadow duration-300">
            <div className="flex items-center p-6">
              <div className="w-12 h-12 rounded-full bg-red-100 flex items-center justify-center mr-4">
                <svg
                  className="w-6 h-6 text-red-500"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M13 10V3L4 14h7v7l9-11h-7z"
                  />
                </svg>
              </div>
              <div>
                <p className="text-sm text-gray-600">Max Temperature</p>
                <h3 className="text-2xl font-bold text-gray-800">
                  {maxTemps[maxTemps.length - 1]}°C
                </h3>
              </div>
            </div>
          </div>

          {/* Min Temperature Card */}
          <div className="bg-white/70 backdrop-blur-sm rounded-xl shadow-lg hover:shadow-xl transition-shadow duration-300">
            <div className="flex items-center p-6">
              <div className="w-12 h-12 rounded-full bg-blue-100 flex items-center justify-center mr-4">
                <svg
                  className="w-6 h-6 text-blue-500"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M20 4v16H4V4M9 4v16M15 4v16"
                  />
                </svg>
              </div>
              <div>
                <p className="text-sm text-gray-600">Min Temperature</p>
                <h3 className="text-2xl font-bold text-gray-800">
                  {minTemps[minTemps.length - 1]}°C
                </h3>
              </div>
            </div>
          </div>

          {/* Avg Temperature Card */}
          <div className="bg-white/70 backdrop-blur-sm rounded-xl shadow-lg hover:shadow-xl transition-shadow duration-300">
            <div className="flex items-center p-6">
              <div className="w-12 h-12 rounded-full bg-yellow-100 flex items-center justify-center mr-4">
                <svg
                  className="w-6 h-6 text-yellow-500"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707"
                  />
                </svg>
              </div>
              <div>
                <p className="text-sm text-gray-600">Avg Temperature</p>
                <h3 className="text-2xl font-bold text-gray-800">
                  {avgTemps[avgTemps.length - 1]}°C
                </h3>
              </div>
            </div>
          </div>

          {/* Alerts Card */}
          <div className="bg-white/70 backdrop-blur-sm rounded-xl shadow-lg hover:shadow-xl transition-shadow duration-300">
            <div className="flex items-center p-6">
              <div className="w-12 h-12 rounded-full bg-orange-100 flex items-center justify-center mr-4">
                <svg
                  className="w-6 h-6 text-orange-500"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
                  />
                </svg>
              </div>
              <div>
                <p className="text-sm text-gray-600">Total Alerts</p>
                <h3 className="text-2xl font-bold text-gray-800">
                  {Object.values(alertsPerCity).reduce((a, b) => a + b, 0)}
                </h3>
              </div>
            </div>
          </div>
        </div>

        {/* Main Charts Grid */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          {/* Temperature Trends Chart */}
          <div className="bg-white/70 backdrop-blur-sm rounded-xl shadow-lg p-6 hover:shadow-xl transition-shadow duration-300">
            <h2 className="text-lg font-semibold text-gray-800 mb-4">
              Temperature Trends (Last 7 Days)
            </h2>
            <div className="p-4">
              <Line
                data={lineChartData}
                options={{
                  responsive: true,
                  plugins: {
                    legend: {
                      position: 'bottom',
                    },
                  },
                  scales: {
                    y: {
                      beginAtZero: false,
                      grid: {
                        color: 'rgba(0,0,0,0.05)',
                      },
                    },
                    x: {
                      grid: {
                        display: false,
                      },
                    },
                  },
                }}
              />
            </div>
          </div>

          <div className="bg-white/70 backdrop-blur-sm rounded-xl shadow-lg p-6 hover:shadow-xl transition-shadow duration-300">
            <h2 className="text-lg font-semibold text-gray-800 mb-4">
              Average Temperature Trend
            </h2>
            <div className="p-4">
              <Bar
                data={barChartData}
                options={{
                  responsive: true,
                  plugins: {
                    legend: {
                      position: 'bottom',
                    },
                  },
                  scales: {
                    y: {
                      beginAtZero: false,
                      grid: {
                        color: 'rgba(0,0,0,0.05)',
                      },
                    },
                    x: {
                      grid: {
                        display: false,
                      },
                    },
                  },
                }}
              />
            </div>
          </div>

          {/* Weather Conditions Chart */}
          <div className="bg-white/70 backdrop-blur-sm rounded-xl shadow-lg p-6 hover:shadow-xl transition-shadow duration-300">
            <h2 className="text-lg font-semibold text-gray-800 mb-4">
              Weather Conditions Distribution
            </h2>
            <div className="p-4">
              <Pie
                data={pieChartData}
                options={{
                  responsive: true,
                  plugins: {
                    legend: {
                      position: 'bottom',
                    },
                  },
                }}
              />
            </div>
          </div>

          {/* Alerts by City Chart */}
          <div className="bg-white/70 backdrop-blur-sm rounded-xl shadow-lg p-6 lg:col-span-2 hover:shadow-xl transition-shadow duration-300">
            <h2 className="text-lg font-semibold text-gray-800 mb-4">
              Alerts by City
            </h2>
            <div className="p-4">
              <Bar
                data={alertsBarChartData}
                options={{
                  ...alertsBarChartOptions,
                  responsive: true,
                  plugins: {
                    legend: {
                      display: false,
                    },
                  },
                  scales: {
                    x: {
                      grid: {
                        display: false,
                      },
                    },
                    y: {
                      grid: {
                        display: false,
                      },
                    },
                  },
                }}
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default ChartsPage
