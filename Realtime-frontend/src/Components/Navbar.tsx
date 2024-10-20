import { Link } from 'react-router-dom'

const Navbar = () => {
  return (
    <nav className="bg-blue-600 p-4">
      <div className="max-w-3xl mx-auto flex justify-between items-center">
        <Link to="/config" className="text-white hover:text-blue-200 ">
          Config
        </Link>
        <Link to="/dashboard" className="text-white hover:text-blue-200">
          Charts
        </Link>
      </div>
    </nav>
  )
}
export default Navbar
