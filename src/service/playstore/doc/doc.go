package doc

/*
google  没有想sb苹果一样提供一个查询接口，google 采用的数字签名的方式
detail:
http://developer.android.com/google/play/billing/billing_reference.html

The getBuyIntent() method

This method returns a response code integer mapped to the RESPONSE_CODE key, and a PendingIntent to launch the purchase flow for the in-app item mapped to the BUY_INTENT key, as described in Purchasing an Item. When it receives the PendingIntent, Google Play sends a response Intent with the data for that purchase order. The data that is returned in the response Intent is summarized in table 3.

Table 3. Response data from an In-app Billing Version 3 purchase request.
keyDescription
RESPONSE_CODEValue is 0 if the purchase was success, error otherwise.
INAPP_PURCHASE_DATAA String in JSON format that contains details about the purchase order. See table 4 for a description of the JSON fields.
INAPP_DATA_SIGNATURE String containing the signature of the purchase data that was signed with the private key of the developer. The data signature uses the RSASSA-PKCS1-v1_5 scheme.
*/
