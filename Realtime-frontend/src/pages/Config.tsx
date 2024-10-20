import { ChangeEvent, FormEvent, useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'

function ConfigPage(): JSX.Element {
  const [cities, setCities] = useState<string>('')
  const [thresholdTemp, setThresholdTemp] = useState<number>()
  const [email, setEmail] = useState<string>()
  const [isConfigSaved, setIsConfigSaved] = useState<boolean>(false)
  const [consecutiveAlertCount, setConsecutiveAlertCount] = useState<number>()
  const navigate = useNavigate()

  const handleCitiesChange = (
    event: ChangeEvent<HTMLTextAreaElement>
  ): void => {
    setCities(event.target.value)
  }

  useEffect(() => {
    const fetchConfig = async () => {
      const res = await fetch(`${import.meta.env.VITE_SERVER_URL}/config`)
      const data = await res.json()
      if (!res.ok) {
        console.log('No config found')
        return
      }
      setIsConfigSaved(true)
      setCities(data.cities.join(', '))
      //set in local storage
      localStorage.setItem('cities', JSON.stringify(data.cities))
      setThresholdTemp(data.threshold_temperature)
      setEmail(data.email)
      setConsecutiveAlertCount(data.consecutive_alert_threshold)
    }
    fetchConfig()
  }, [])

  const handleSubmit = async (
    event: FormEvent<HTMLFormElement>
  ): Promise<void> => {
    event.preventDefault()
    // Save the configuration (e.g., send to backend or store locally)
    console.log({
      cities,
      thresholdTemp,
      email,
      consecutiveAlertCount,
    })

    const cityList = cities.split(',').map((city) => city.trim().toLowerCase())

    const res = await fetch(`${import.meta.env.VITE_SERVER_URL}/config`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        cities: cityList,
        threshold_temperature: thresholdTemp,
        email,
        consecutive_alert_threshold: consecutiveAlertCount,
      }),
    })

    if (!res.ok) {
      alert('Error saving configuration')
      console.log(res.status)
      return
    }

    // Navigate to Charts page
    navigate('/dashboard')
  }

  if (isConfigSaved) {
    return (
      <>
        <div className="max-w-3xl mx-auto my-8 p-8 bg-gradient-to-br from-blue-50 to-indigo-100 rounded-xl shadow-lg">
          <h2 className="text-3xl font-bold mb-8 text-indigo-800 border-b-2 border-indigo-200 pb-2">
            Current Configuration
          </h2>
          <div className="space-y-6">
            <div className="bg-white p-4 rounded-lg shadow-md">
              <h3 className="text-lg font-semibold text-indigo-600 mb-2">
                Selected Cities
              </h3>
              <div className="flex flex-wrap gap-2">
                {cities.split(',').map((city) => (
                  <span
                    key={city}
                    className="px-3 py-2 bg-indigo-100 text-indigo-800 rounded-full text-sm"
                  >
                    {city}
                  </span>
                ))}
              </div>
            </div>
            <div className="bg-white p-4 rounded-lg shadow-md">
              <h3 className="text-lg font-semibold text-indigo-600 mb-2">
                Threshold Temperature
              </h3>
              <p className="text-2xl font-bold text-gray-800">
                {thresholdTemp}°C
              </p>
            </div>
            <div className="bg-white p-4 rounded-lg shadow-md">
              <h3 className="text-lg font-semibold text-indigo-600 mb-2">
                Email
              </h3>
              <p className="text-gray-800">{email}</p>
            </div>
            <div className="bg-white p-4 rounded-lg shadow-md">
              <h3 className="text-lg font-semibold text-indigo-600 mb-2">
                Consecutive Alert Count Threshold
              </h3>
              <p className="text-2xl font-bold text-gray-800">
                {consecutiveAlertCount}
              </p>
            </div>
          </div>
        </div>
      </>
    )
  }

  return (
    <div className="max-w-3xl mx-auto my-8 p-6">
      <h2 className="text-2xl font-bold mb-6">Configuration</h2>
      <form onSubmit={handleSubmit} className="space-y-6">
        {/* Cities Selection */}
        <div>
          <label className="block text-gray-700 font-medium mb-2">
            Enter Cities (comma-separated):
          </label>
          <textarea
            value={cities}
            onChange={handleCitiesChange}
            required
            className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition duration-200 ease-in-out resize-none h-30 bg-white shadow-inner text-gray-700 placeholder-gray-400"
            placeholder="e.g. Chennai, Mumbai, Kolkata, Delhi"
          />
        </div>
        {/* Threshold Temperature */}
        <div>
          <label className="block text-gray-700 font-medium mb-2">
            Threshold Temperature (°C):
          </label>
          <input
            type="number"
            value={thresholdTemp}
            onChange={(e: ChangeEvent<HTMLInputElement>) =>
              setThresholdTemp(e.target.value as unknown as number)
            }
            required
            className="w-full px-3 py-2 border border-gray-300 rounded-md"
          />
        </div>
        {/* Email */}
        <div>
          <label className="block text-gray-700 font-medium mb-2">Email:</label>
          <input
            type="email"
            value={email}
            onChange={(e: ChangeEvent<HTMLInputElement>) =>
              setEmail(e.target.value)
            }
            required
            className="w-full px-3 py-2 border border-gray-300 rounded-md"
          />
        </div>
        {/* Consecutive Alert Count Threshold */}
        <div>
          <label className="block text-gray-700 font-medium mb-2">
            Consecutive Alert Count Threshold:
          </label>
          <input
            type="number"
            value={consecutiveAlertCount}
            onChange={(e: ChangeEvent<HTMLInputElement>) =>
              setConsecutiveAlertCount(e.target.value as unknown as number)
            }
            required
            className="w-full px-3 py-2 border border-gray-300 rounded-md"
          />
        </div>
        {/* Submit Button */}
        <button
          type="submit"
          className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700"
        >
          Save Configuration
        </button>
      </form>
    </div>
  )
}

export default ConfigPage
