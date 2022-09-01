import React from 'react'
import { MainLayout } from '../../src/components/Layout/MainLayout'
import { Elements, useElements } from '@stripe/react-stripe-js'
import { loadStripe } from '@stripe/stripe-js'
import { PaymentElement } from '@stripe/react-stripe-js'

const stripePromise = loadStripe(process.env.public_stripe_key)

const index = () => {
  const elements = useElements()

  // //クレカの取り扱い
  // const token = await Stripe.createToken

  return (
    <Elements stripe={stripePromise}>
      <MainLayout title="stripe">
        <form>
          <PaymentElement />
          <button>Submit</button>
        </form>
      </MainLayout>
    </Elements>
  )
}

export default index
