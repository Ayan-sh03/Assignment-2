import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { createBrowserRouter, RouterProvider, Outlet } from 'react-router-dom'
import ConfigPage from './pages/Config.tsx'
import ChartsPage from './pages/Charts.tsx'
import Navbar from './Components/Navbar.tsx'

// Create a layout component that includes the Navbar
const Layout = () => {
  return (
    <>
      <Navbar />
      <Outlet />
    </>
  )
}

const router = createBrowserRouter([
  {
    element: <Layout />,
    children: [
      {
        path: '/',
        element: <App />,
      },
      {
        path: '/config',
        element: <ConfigPage />,
      },
      {
        path: '/dashboard',
        element: <ChartsPage />,
      },
    ],
  },
])

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>
)
