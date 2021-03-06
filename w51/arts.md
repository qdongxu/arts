# Algorithm

  Shortest path(Dijkstra) 

# Reading

    [https://blog.cloudflare.com/how-to-receive-a-million-packets/]

# Tips
  N/A

# Sharing

## Zuul Introduction
 
### Backgroud

 - Netflix was migrating to public cloud in 2007. 
 - Micro-services became the major architecture 
 
There were serval components or SDK created:
 - Ribbon: client-side load balance
 - Eureka: service registration/descovery
 - Hystrix: falt tolerance
 - Archaius: configuration API
 - Zuul:  API gateway
 

Open sourced around 2012. 

### Function Overview

- Authentication

- Insights and Monitoring 
	- tracking and statistic at the edge
	
- Routing 
	- dynamically routing 
	
- Stress Testing 
	- increasing the traffic to a cluster 
	
- Load Shedding 
	- Rate Limit

- Static Response handling 
	- Fallback
	- Cacheing

- Multiregion Resiliency 
	- Cross IDC
	- CDN

### Architecture Overview
#### Main Structure
<div align="center"><img src="https://miro.medium.com/max/1266/1*j9iGkeQ7bPK2nC1a7BgFOw.png" width="75%"></img></div>

- Request Context
- Filter Runner
- Filter (Groovy)
- Filter Loader (Dynamicaly)

#### Request Process:
 
<div align="center"><img src="https://miro.medium.com/max/1200/1*9IeEGHSRMGfAnhqM49TLpQ.png" width="75%"></img></div>

**Note:** A special pathway: 
 - Error filter

### Nonfunctional Feature

1. Performance

	Zuul 1:
    - 800 ~ 1000 tps 
    - (30m Nginx) ( 200ms Zuul 1 )
    - bad 99% response time 
    
    <br/>
	zuul 2( Dec 2017)
	
	<div align="center"><img src="https://miro.medium.com/max/1495/0*KUIIWlngbnDUDvLA." width="75%"></img></div>
	Request per Second
	<br/>
	<div align="center"><img src="https://miro.medium.com/max/1484/0*nBUn2N0zyQO1hxbB." width="75%"></img></div>

2. API oriented
	- interactive with other services
3. Typical Scenario
	- filter written by dynamic language, groovy
	- upload a temporary filter to route some special request to specific node for trouble shooting.
	 
### Zuul 2

1. large scale of long connections
2. high performance
	- CPU decreased 25% while throughput increased 25% , ( by the condition that less filters) (Sep 2016)

### Way Forward

1. Netflix will probably port zuul 2  to Spring Cloud Gateway. 

### Reference




