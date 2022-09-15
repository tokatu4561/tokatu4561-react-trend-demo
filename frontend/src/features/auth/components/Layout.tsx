import Head from 'next/head'
import React from 'react'

interface LayoutProps {
  children: React.ReactNode
  title: string
}

export const Layout = ({ children, title }: LayoutProps) => {
  return (
    <>
      <Head>
        <title>{title}</title>
      </Head>
      <div className="bg-amber-50 h-screen flex">
        <div className="flex-1">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
            {children}
          </div>
        </div>
      </div>
    </>
  )
}
