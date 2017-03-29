package apiresp

//GetOrderAggRequest Function to generate Order Aggregator Request body
func GetOrderAggRequest(orderNumber string) string {
	return `{
    order(id: "` + orderNumber + `") {
      id
      lock {
        userId
        location
        id
      }
      created
      currentTotal {
        amount
        currency
      }
      originalTotal {
        amount
        currency
      }
      returnTotal {
        amount
        currency
      }
      netTotal {
        amount
        currency
      }
      promos {
        description
        longDescription
        promoType
        totalDiscount {
          currency
          amount
        }
        _promo_id
      }
      billee {
        firstName
        lastName
        dayPhone
        emailAddress
      }
      paymentMethods {
        paymentType
      }
      services {
        _service_id
        charges {
          baseCharge {
            currency
            amount
          }
          serviceCharges {
            currency
            amount
          }
          reservationWindowCharges {
            currency
            amount
          }
        }
      }
      lines {
        hostOrderLineReference
        type {
          description
          code
        }
        poNumber
        quantityOrdered
        originalOrderedQty
        adjustments {
          currency
          amount
              category
            _promo_id
            name
          }
        adjustmentTotals {
          discount
          charge
          shipDiscount
          shipCharge
          netChargeLessShip
          }
        item {
          omsId
          sku
          upc
          vendor {
            name
            number
          }
          description
          shortDescription
          blindsConfiguration {
            configurationId
            leadTime
          }
          unitCost {
            currency
            amount
          }
          unitPrice {
            currency
            amount
          }
          unitWeight {
            amount
            units
          }
          units
          class
          subclass
          department
        }
        fulfillment {
          description
          code
        }
        statuses {
          quantity
          description
          code
        }
        statusesDetails {
          quantity
          description
          code
        }
        tax {
          amount
          currency
          taxName
        }
        appliedPromos {
          discount {
            currency
            amount
          }
          _promo_id
        }
        pickUp {
          desk
          deliveryHasBeenNotified
          quantityPicked
          quantityReceived
          recipient {
            email
            firstName
            lastName
            phone
          }
          estimatedEarliestDate
          estimatedLatestDate
          store {
            phone
            name
            number
          }
          status
        }
        deliveries {
          workOrderNo
          service
          routeId
          storeId
          _service_id
          isUnattended
          recipient {
            firstName
            lastName
            phone
            address{
              line1
              line2
              line3
              line4
              line5
              city
              state
              zipCode
              country
            }
          }
          reservationWindow {
            end
            start
          }

          instructions {
            sequenceNo
            text
          }
          charge {
            amount
            currency
          }
          trackingStatus {
            description
            code
          }
        }
        shipping {
          shipmentModeCode
          origin
          service {
            description
            code
          }
          recipient {
            emailAddress
            title
            suffix
            firstName
            middleName
            lastName
            company
            dayPhone
            mobilePhone
            otherPhone
            address{
              line1
              line2
              line3
              line4
              line5
              city
              state
              zipCode
              country
          }
          }
          discounts {
            currency
            amount
          }
          charges{
            currency
            amount
          }
          estimatedDeliveryStartDate
          estimatedDeliveryEndDate
          shipments {
            trackingNumbers
            provider
        }
        }
        tags
      }
      id
      }
  }`
}
