import React from 'react'
import { CardElement, useElements, useStripe } from '@stripe/react-stripe-js'
import axios from 'axios'

export const CheckoutForm = () => {
  const stripe = useStripe()
  const elements = useElements()

  // //クレカの取り扱い
  // const token = await Stripe.createToken

  const handleFormSubmit = async (event) => {
    const response = await axios.post('/api/pay', {
      body: JSON.stringify({ amount: event.target.amount.value }),
    })

    const data = response.data
    const clientSecret = data.client_secret

    const card = elements.getElement(CardElement)

    const result = await stripe.confirmCardPayment(clientSecret, {
      payment_method: {
        card: card,
        billing_details: {
          name: 'test',
        },
      },
    })

    if (result.error) {
      // card declined, or something went wrong with the card
      console.log(result.error)
    } else if (result.paymentIntent) {
      if (result.paymentIntent.status === 'succeeded') {
        console.log('success')
        // we have charged the card
      }
    }
  }

  return (
    <form onSubmit={handleFormSubmit}>
      {/* <PaymentElement /> */}

      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="username"
        >
          Username
        </label>
        <input
          className="shadow border rounded w-full py-2 px-3 text-gray-700"
          type="text"
          placeholder="Username"
        />
      </div>
      <div className="mb-4">
        <CardElement />
      </div>
      <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
        決済する
      </button>
    </form>
  )
}
