---
title: Architecture Overview
weight: 0
draft: true
publishDate: 2999-12-31
---
# Architecture Overview

The following documentation provides an overview of the system.

## Table of Contents

- [Introduction and Goals](01_introduction_and_goals.md)
    Short description of the requirements, driving forces, extract (or abstract) of requirements. Top three (max five) quality goals for the architecture which have highest priority for the major stakeholders. A table of important stakeholders with their expectation regarding architecture.
- [Architecture Constraints](02_architecture_constraints.md)
    Anything that constrains teams in design and implementation decisions or decision about related processes. Can sometimes go beyond individual systems and are valid for whole organizations and companies.

- [System Scope and Context](03_system_scope_and_context.md)
    Delimits the system from its (external) communication partners (neighboring systems and users). Specifies the external interfaces. Shown from a business/domain perspective (always) or a technical perspective (optional)

- [Solution Strategy](04_solution_strategy.md)
    Summary of the fundamental decisions and solution strategies that shape the architecture. Can include technology, top-level decomposition, approaches to achieve top quality goals and relevant organizational decisions.

- [Building Blocks](05_building_block_view.md)
    Static decomposition of the system, abstractions of source-code, shown as hierarchy of white boxes (containing black boxes), up to the appropriate level of detail.

- [Runtime View](06_runtime_view.md)
    Behavior of building blocks as scenarios, covering important use cases or features, interactions at critical external interfaces, operation and administration plus error and exception behavior.

- [Deployment View](07_deployment_view.md)
    Technical infrastructure with environments, computers, processors, topologies. Mapping of (software) building blocks to infrastructure elements.

- [Concepts](08_concepts.md)
    Overall, principal regulations and solution approaches relevant in multiple parts (→ cross-cutting) of the system. Concepts are often related to multiple building blocks. Include different topics like domain models, architecture patterns and -styles, rules for using specific technology and implementation rules.

- [Design Descisions](09_design_decisions.md)
    Important, expensive, critical, large scale or risky architecture decisions including rationales.

- [Quality Scenarios](10_quality_scenarios.md)
    Quality requirements as scenarios, with quality tree to provide high-level overview. The most important quality goals should have been described in section 1.2. (quality goals).

- [Technical Risks](11_technical_risks.md)
    Known technical risks or technical debt. What potential problems exist within or around the system? What does the development team feel miserable about?

- [Glossary](12_glossary.md)
    Important domain and technical terms that stakeholders use when discussing the system. Also: translation reference if you work in a multi-language environment.

