Plotly Go API
=

This package provides wrapper functions for [Plotly](https://plot.ly)'s HTTP API.

At the time of writing that API is largely undocumented so the work here is based on
the [Pyhton API](https://github.com/plotly/python-api).

Authentication
==

In order to use this package you require API credentials from Plotly. These may
be stored in:

1. `.plotly_credentials.json`
2. `plotly_credentials.json`
3. `/etc/plotly/.plotly_credentials.json`
4. `/etc/plotly/plotly_credentials.json`
5. Environment variables named `PLOTLY_USERNAME` and `PLOTLY_APIKEY`

If more than one of these are available, the highest one in the list takes preference.

The `.json` files should contain the following:

    {"Username":"yourname","Apikey":"yourkey"}


Usage
==

An example program is provided in this repository.

Limitations
==

This is a work in process.

One important thing to be aware of is that the plotly API always returns 200,
so checking for an error from the request is not suitable, rather look at the
`error` field in the response.
