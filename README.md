# trackid

[![Build Status](https://travis-ci.org/go-guoyk/trackid.svg?branch=master)](https://travis-ci.org/go-guoyk/trackid)

general track id for context.Context

## Usage

**Set TrackId**

```go
ctx = trackid.Set(ctx, "111")
```

**Get TrackId**

```go
trackid.Get(ctx)
```

**Clear TrackId**

```go
trackid.Clear(ctx)
```

## Credits

Guo Y.K. <hi@guoyk.net>
