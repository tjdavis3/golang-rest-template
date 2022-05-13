---
title: Cross-cutting Concepts
linkTitle: Concepts
description: Describes overall, principal regulations and solution ideas that are relevant in multiple parts of the system
weight: 8
draft: false
---

<!--

**Content.**

This section describes overall, principal regulations and solution ideas
that are relevant in multiple parts (= cross-cutting) of your system.
Such concepts are often related to multiple building blocks. They can
include many different topics, such as

  - domain models

  - architecture patterns or design patterns

  - rules for using specific technology

  - principal, often technical decisions of overall decisions

  - implementation rules

**Motivation.**

Concepts form the basis for *conceptual integrity* (consistency,
homogeneity) of the architecture. Thus, they are an important
contribution to achieve inner qualities of your system.

Some of these concepts cannot be assigned to individual building blocks
(e.g. security or safety). This is the place in the template that we
provided for a cohesive specification of such concepts.

**Form.**

The form can be varied:

  - concept papers with any kind of structure

  - cross-cutting model excerpts or scenarios using notations of the
    architecture views

  - sample implementations, especially for technical concepts

  - reference to typical usage of standard frameworks (e.g. using
    Hibernate for object/relational mapping)

**Structure.**

A potential (but not mandatory) structure for this section could be:

  - Domain concepts

  - User Experience concepts (UX)

  - Safety and security concepts

  - Architecture and design patterns

  - "Under-the-hood"

  - development concepts

  - operational concepts

Note: it might be difficult to assign individual concepts to one
specific topic on this list.

![Possible topics for crosscutting
concepts](images/08-Crosscutting-Concepts-Structure-EN.png)

-->

## Domain Concepts

* 

* 

* 

* 

* 

* 

## User-Experience (UX)

```plantuml
@startsalt
{+
{#
**Cust Circuit** | **Cust** | **CID** | **Alt CID** | **Vendor**
M5-xxxxxx-xxxxxx | 138 | 52/KGFN/110606/TWCS |. | Level 3
}
{
[ <&check> OK ]  | [ Cancel ]
}
}
@endsalt
```

## Operational Concepts

```plantuml
@startuml

!include <c4/C4_Context.puml>
!include <office/Users/user.puml>
!include <office/Users/approver.puml>
!include <logos/github-icon.puml>
!include <logos/prometheus.puml>
!include <logos/grafana.puml>
!include <logos/sentry.puml>
!include <office/Users/writer.puml>
!include <office/Users/administrator.puml>
!include <logos/drone.puml>
!include <logos/github-icon>
!include <logos/jaeger>
!include <material/document.puml>

title Application and Service Features

System(app, RingSquared Services)
Person(dev, "Developer", ,writer)
Person(ops, "Operations", ,administrator)
Person(user, "End-user", )

System_Ext(prom, Prometheus, Gathers metrics and alerts on rules violations, prometheus)
System_Ext(grafana, "Grafana", "Provides visualization dashboards of metrics, tracing, and logs",grafana)
System_Ext(tempo, Tempo, Distributed Tracing, jaeger)
System_Ext(loki, "Loki", Log Aggregation,ma_document)
System_Ext(sentry, Sentry, Tracks system exceptions,sentry)
System_Ext(alert, AlertManager, Manages alerts from prometheus rules, prometheus)
System_Ext(ch, "Github", "Tracks all development work and projects", github-icon)

Rel_U(prom, app, "Pulls application metrics")
Rel_U(grafana, prom, "Pulls metrics to display on dashboard")
Rel_U(grafana, tempo, Pulls tracing information)
Rel_U(grafana, loki, "Pulls logs for display on dashboard")
Rel_R(grafana, alert, "Pulls alert details")
Rel_D(app, sentry, "Pushes errors and crash reports")
Rel_D(app, loki, "Sends log messages")
Rel_D(app, tempo, "Sends tracing spans")
Rel_D(sentry, ch, "Creates stories for unresolved errors")
Rel(prom, alert, Triggers alerts)
Rel(alert, ops, Sends alerts)

Rel_U(dev, ch, "Pulls stories to fix")
Rel_U(ops, grafana, "Monitors dashboards for system performance")
Rel_R(user, app, "Uses the application")

footer All applications are being built with these features integrated

@enduml
```

The service interacts with several external systems to provide observability into the performance of the system.

**Loki**
: Centralized logging system from Grafana.  The application logs to stdout and the underlying operating system then uses promtail to forward the logs to Loki.  Promtail can be run standalone on servers or as a sidecar for kubernetes and docker containers.

**Tempo**
: Distributed tracing system from Grafana.  This system is able to receive traces written for various telemtry systems including OpenTelemetry, Zipkin, and Jaeger.  We use OpenTelemetry to send the traces and spans.

**Prometheus**
: Metrics gathering system.  Each service is written to expose metrics pertaining to the service.  Nodes, databases, kubernetes, queing systems, etc also expose metrics that are gathered by Prometheus.  Alerting rules can be defined based on the collected metrics.

**AlertManager**
: Alerts triggered in Prometheus are forwaded to AlertManager where they can be deduplicated, queued,  and evaluated sending notifications.  AlertManager has the ability to send notifications via Slack, Teams, SMS, email, and a variety of other methods.

**Grafana**
: Grafana provides a centralized place to view dashboards that correlate metrics, logs, traces, and alerts.  

**Sentry**
: Crashes and unexpected errors are forwarded to Sentry.  There they are collected, alerts sent out, and issues created in Github for further exploration.


