# New Relic Go Agent

An unofficial [New Relic][nr] agent for the Go programming language. Allows your Go web
app to appear as an application in New Relic (under the "APM" tab).

It wraps the [New Relic Agent SDK][sdk] -- the SDK is currently Linux only, therefore,
so is this Go agent.

Features:
 
* Transaction monitoring -- from top of a web request to bottom
* Segments -- external calls, database calls, etc.
* Custom metrics

[nr]: https://newrelic.com/
[sdk]: https://docs.newrelic.com/docs/agents/agent-sdk/using-agent-sdk/getting-started-agent-sdk
