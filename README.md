# Automation Process for Creating a Twilio Subaccount and Getting A2P 10DLC Verified

This guide outlines the steps necessary to automate the creation of a Twilio subaccount and the process of getting A2P 10DLC verified.

## 1. Create a Secondary Customer Profile
- **Description:** Create a secondary customer profile and attach the necessary business information and authorized representatives to the profile.
- **Steps:**
  1. Fetch available `PolicySID` values.
  2. Use the selected `PolicySID` to create the customer profile.
  3. Attach business information to the profile.
  4. Add authorized representatives to the profile.

## 2. Create and Submit a TrustProduct
- **Description:** Create a TrustProduct, attach the necessary resources, evaluate, and submit it for review.
- **Steps:**
  1. Create a TrustProduct resource.
  2. Attach the required resources, such as customer profiles and other necessary documents.
  3. Evaluate the TrustProduct to ensure all required fields are correctly filled.
  4. Submit the TrustProduct for review.

## 3. Create a BrandRegistration
- **Description:** Register the customer's brand with The Campaign Registry (TCR) for vetting.
- **Steps:**
  1. Provide brand details, including company name, address, and other required information.
  2. Submit the brand registration for vetting by TCR.

## 4. Create a MessagingService
- **Description:** Set up a messaging service to handle A2P 10DLC messaging.
- **Steps:**
  1. Create a new messaging service.
  2. Configure the messaging service to use the desired features such as sticky sender, fallback, or smart encoding.

## 5. Create an A2P Campaign
- **Description:** Define the campaign, including the use case and messaging flow.
- **Steps:**
  1. Create a campaign for the A2P 10DLC service.
  2. Specify the use case, message samples, and flow of messages.

## 6. Add a Phone Number to the Messaging Service
- **Description:** Associate a phone number with the messaging service to enable message sending.
- **Steps:**
  1. Purchase a phone number if not already available.
  2. Add the purchased phone number to the messaging service.

## 7. Submit the Campaign for Review
- **Description:** Finalize the campaign setup and submit it for approval.
- **Steps:**
  1. Review the campaign details and ensure everything is in order.
  2. Submit the campaign for final approval by Twilio.

## 8. Finalize the Setup, Monitor, and Troubleshoot
- **Description:** Finalize the overall setup and continuously monitor the campaign and messaging services. Troubleshoot any issues that arise during the process.
- **Steps:**
  1. Verify that the campaign is live and working as expected.
  2. Monitor the messaging service for delivery status and error codes.
  3. Address any issues promptly to maintain compliance and messaging quality.

## 9. Fetching and Using PolicySID
- **Description:** Fetch available policies and use the selected `PolicySID` for creating customer profiles and trust products.
- **Steps:**
  1. Retrieve the list of available policies from Twilio.
  2. Use the appropriate `PolicySID` in customer profiles and TrustProduct creation.

## 10. Purchasing and Associating Phone Numbers
- **Description:** Purchase a phone number and associate it with a messaging service to enable your customers to send messages using A2P 10DLC.
- **Steps:**
  1. Purchase a phone number from Twilio.
  2. Add the phone number to the messaging service as per the requirements of the A2P 10DLC campaign.

## Conclusion
This document covers the full process for A2P 10DLC registration and messaging setup in Twilio for your customers. By following these steps, you can automate the creation of a Twilio subaccount and successfully register for A2P 10DLC, ensuring compliance and effective messaging.

---

## This Custom Twilio-Go Package is Created by **[Vikash Parashar](https://www.linkedin.com/in/vikash-parashar-3152471ba)**
## Resources Link's
1. [Take a pull from github](https://github.com/vikash-parashar/twl-a2p-auto)
