# Disposable Email Checker

Simple service that helps you check disposable emails.

## Usage

Make a GET request to the `/check/:email` endpoint, where `:email` is an email
you want to check. If the email is disposable email address you get the
following response:

```json
{
  "disposable": true,
  "success": true
}
```

**Note:** The [list of disposable email domains](https://github.com/disposable/disposable-email-domains) is automatically updated every 24
hours.

## Deployment

We are using Disposable Email Checker internally, and **we are not** providing
public API. You can deploy this service on [Railway](https://railway.app) with
one-click using the button below.

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/new/template/tHOg9a?referralCode=saliven)

If you want to change the list, you can change the value of the `DOMAIN_LIST`
environment variable.
