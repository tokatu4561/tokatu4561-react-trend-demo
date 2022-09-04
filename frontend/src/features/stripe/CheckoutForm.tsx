import React from 'react'
import { CardElement, useElements, useStripe } from '@stripe/react-stripe-js'
import { InputFiled } from '../../components/Form/InputFiled'
import { axios } from '@/lib/axios'

export const CheckoutForm = () => {
  const stripe = useStripe()
  const elements = useElements()

  // //クレカの取り扱い
  // const token = await Stripe.createToken

  const handleFormSubmit = async (event) => {
    event.preventDefault()
    const { data } = await axios.post('/api/payment-intent', {
      amount: '1000',
    })

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
      console.log(result.error)
    } else if (result.paymentIntent) {
      if (result.paymentIntent.status === 'succeeded') {
        console.log('success')
      }
    }
  }

  return (
    <form onSubmit={handleFormSubmit}>
      <div className="mb-4">
        <InputFiled
          inputId="username"
          label="Username"
          type="text"
          placeholder="Username"
        />
      </div>
      <div className="mb-4">
        <InputFiled
          inputId="amount"
          label="Amount"
          type="text"
          placeholder="amount"
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
