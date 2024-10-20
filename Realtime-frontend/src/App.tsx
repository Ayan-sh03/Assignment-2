import './App.css'

function App() {
  return (
    <div className="min-h-screen bg-gradient-to-r from-blue-400 to-purple-500 flex items-center justify-center">
      <div className="bg-white p-8 rounded-3xl shadow-lg text-center">
        <h1 className="text-4xl font-bold text-gray-800 mb-4">
          Welcome to Our App!
        </h1>
        <p className="text-xl text-gray-600 mb-6">
          Experience the future of technology.
        </p>
        <a
          href="/config"
          className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-full transition duration-300"
        >
          Get Started
        </a>
      </div>
    </div>
  )
}

export default App
