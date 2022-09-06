import Head from 'next/head'
import Link from 'next/link'
import * as React from 'react'

type Props = {
  children: React.ReactNode
  title: string
}

export const MainLayout = ({ children, title }: Props) => {
  return (
    <>
      <Head>
        <title>{title}</title>
      </Head>
      <div className="bg-amber-50 h-screen flex">
        <SideBar />
        <div className="flex-1">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
            {children}
          </div>
        </div>
      </div>
    </>
  )
}

const SideBar = () => {
  return (
    <aside className="w-64" aria-label="Sidebar">
      <div className="overflow-y-auto py-4 px-3 h-full bg-gray-900 rounded">
        <ul className="space-y-2">
          {SideBarNavigationItem.map((navItem) => (
            <li key={navItem.name}>
              <Link href={navItem.href}>
                <a className="flex items-center p-2 text-base font-normal text-white rounded-lg hover:bg-gray-100">
                  <span className="ml-3">{navItem.name}</span>
                </a>
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </aside>
  )
}

type SideBarNavigationItemType = {
  name: string
  href: string
}

const SideBarNavigationItem: SideBarNavigationItemType[] = [
  {
    name: 'Home',
    href: '/',
  },
  { name: 'Stripe', href: '/stripe-test' },
  {
    name: 'React Hook Form',
    href: '/form-test/react-hook-form',
  },
  {
    name: 'React Hook Form Test',
    href: '/form-test/react-hook-form/test',
  },
  {
    name: 'Formik Test',
    href: '/form-test/formik/test',
  },
]
