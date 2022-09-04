import React from 'react'
import { MainLayout } from '../../src/components/Layout/MainLayout'
import { Elements } from '@stripe/react-stripe-js'
import { loadStripe } from '@stripe/stripe-js'
import { CheckoutForm } from '../../src/features/stripe/CheckoutForm'

const stripeApiKey = `${process.env.PUBLIC_STRIPE_KEY}`

const StripeTest = () => {
  const stripePromise = loadStripe(stripeApiKey)
  return (
    <MainLayout title="stripe">
      <Elements stripe={stripePromise}>
        <div className="h-screen flex justify-center items-center">
          <div className="flex-[0.6_1_0%]">
            <CheckoutForm />
          </div>
        </div>
      </Elements>
    </MainLayout>
  )
}

export default StripeTest
