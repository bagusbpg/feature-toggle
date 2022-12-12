# About
This repository demonstrates the implementation of simple feature toggle in Golang. It uses a package provided by Unleash, an open-source feature management solution.

# How to
First, we need to start the Unleash Server. The easiest way is by running the following script in your terminal.
```bash
git clone git@github.com:Unleash/unleash-docker.git
cd unleash-docker
docker compose up --detach
```
Open your browser and go to localhost:4242. You need to log into Unleash dashbord using this credential:
```
username : admin
password : unleash4all
```
Later, you will use this dashboard to create and configure your feature toggle.


Back to your terminal, you can initialize your Golang project as usual
```bash
mkdir <your-project-directory>
cd <your-project-directory>
go mod init <your-project-module-name>
```
You also need to install Unleash Client SDK for Golang
```bash
go get github.com/Unleash/unleash-client-go/v3
```
In your service app, you need to create an Unleash client instance, specifying the url of Unleash server the client will connect to and API token. Typically, you may also want to specify some other options as well, like how often the client should hit the Unleash server to refresh the state of each feature toggles. By default, the refresh interval is 15 seconds. In the following example, you can set the refresh interval to 2 seconds only using `WithRefreshInterval` option.
```go
featureToggleClient, err := unleash.NewClient(
	unleash.WithListener(&unleash.DebugListener{}),
	unleash.WithAppName("my-application"),
	unleash.WithUrl("http://localhost:4242/api/"),
	unleash.WithRefreshInterval(2*time.Second),
	unleash.WithCustomHeaders(http.Header{"Authorization": {"default:development.unleash-insecure-api-token"}}),
)
if err != nil {
	panic(err)
}
```
At this point, you need to create the feature toggle itself via Unleash Dashboard. The user interface is intuitive. Just click `New feature toggle` at the upper right, and give your feature toggle a name. A naming convention should be agreed by your organization. But for this example, you may name it simmply as `toggle`. It is also nice to specify the type of feature toggle, you may start with `Operational`.

The newly created feature toggle has two environments, `development` and `production`. To serve as an example, you will mostly deal with `development` environment. Click `Add strategy` to your development environment. There are some options available. For starter, you may try `Standard` first.

Once the client and feature toggle are created, generally it is used as follow.
```go
if featureToggleClient.IsEnabled(<NameOfFeatureToggle>) {
	<NewFlow>
} else {
	<OldFlow>
}
```
# References
- Dive deeper to the available [documentations](https://docs.getunleash.io/) and [articles](https://www.getunleash.io/blog/feature-toggle-best-practices) by Unleash to uncover tips and best practices to implement feature toggle in your project.
- Academic research [paper](https://arxiv.org/pdf/1907.06157.pdf) by Mahdavi-Hezaveh, R. et.al. about the use of feature toggle in industry.