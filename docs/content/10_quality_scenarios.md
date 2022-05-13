---
title: Quality Requirements
description: Contains all quality requirements as quality tree with scenarios.
weight: 10
draft: true
---

<!--

**Content.**

This section contains all quality requirements as quality tree with
scenarios. The most important ones have already been described in
section 1.2. (quality goals)

Here you can also capture quality requirements with lesser priority,
which will not create high risks when they are not fully achieved.

**Motivation.**

Since quality requirements will have a lot of influence on architectural
decisions you should know for every stakeholder what is really important
to them, concrete and measurable.

-->


## Quality Tree
<!--

**Content.**

The quality tree (as defined in ATAM – Architecture Tradeoff Analysis
Method) with quality/evaluation scenarios as leafs.

**Motivation.**

The tree structure with priorities provides an overview for a sometimes
large number of quality requirements.

**Form.**

The quality tree is a high-level overview of the quality goals and
requirements:

  - tree-like refinement of the term "quality". Use "quality" or
    "usefulness" as a root

  - a mind map with quality categories as main branches

In any case the tree should include links to the scenarios of the
following section.
-->

<!--

The product quality model defined in ISO/IEC 25010 comprises the eight quality characteristics shown in the following figure:

![ISO25010 Quality Tree](https://iso25000.com/images/figures/en/iso25010.png)


### Functional Suitability
This characteristic represents the degree to which a product or system provides functions that meet stated and implied needs when used under specified conditions. This characteristic is composed of the following sub-characteristics:

Functional completeness 
: Degree to which the set of functions covers all the specified tasks and user objectives.

Functional correctness 
: Degree to which a product or system provides the correct results with the needed degree of precision.

Functional appropriateness 
: Degree to which the functions facilitate the accomplishment of specified tasks and objectives.

### Performance efficiency
This characteristic represents the performance relative to the amount of resources used under stated conditions. This characteristic is composed of the following sub-characteristics:

Time behaviour
: Degree to which the response and processing times and throughput rates of a product or system, when performing its functions, meet requirements.

Resource utilization 
: Degree to which the amounts and types of resources used by a product or system, when performing its functions, meet requirements.

Capacity 
: Degree to which the maximum limits of a product or system parameter meet requirements.

### Compatibility

Degree to which a product, system or component can exchange information with other products, systems or components, and/or perform its required functions while sharing the same hardware or software environment. This characteristic is composed of the following sub-characteristics:

Co-existence
: Degree to which a product can perform its required functions efficiently while sharing a common environment and resources with other products, without detrimental impact on any other product.

Interoperability 
: Degree to which two or more systems, products or components can exchange information and use the information that has been exchanged.

### Usability
Degree to which a product or system can be used by specified users to achieve specified goals with effectiveness, efficiency and satisfaction in a specified context of use. This characteristic is composed of the following sub-characteristics:

Appropriateness recognizability 
: Degree to which users can recognize whether a product or system is appropriate for their needs.

Learnability 
: Degree to which a product or system can be used by specified users to achieve specified goals of learning to use the product or system with effectiveness, efficiency, freedom from risk and satisfaction in a specified context of use.

Operability 
: Degree to which a product or system has attributes that make it easy to operate and control.
User error protection. Degree to which a system protects users against making errors.

User interface aesthetics 
: Degree to which a user interface enables pleasing and satisfying interaction for the user.

Accessibility 
: Degree to which a product or system can be used by people with the widest range of characteristics and capabilities to achieve a specified goal in a specified context of use.

### Reliability
Degree to which a system, product or component performs specified functions under specified conditions for a specified period of time. This characteristic is composed of the following sub-characteristics:

Maturity 
: Degree to which a system, product or component meets needs for reliability under normal operation.

Availability 
: Degree to which a system, product or component is operational and accessible when required for use.

Fault tolerance 
: Degree to which a system, product or component operates as intended despite the presence of hardware or software faults.

Recoverability 
: Degree to which, in the event of an interruption or a failure, a product or system can recover the data directly affected and re-establish the desired state of the system.

### Security
Degree to which a product or system protects information and data so that persons or other products or systems have the degree of data access appropriate to their types and levels of authorization. This characteristic is composed of the following sub-characteristics:

Confidentiality 
: Degree to which a product or system ensures that data are accessible only to those authorized to have access.

Integrity 
: Degree to which a system, product or component prevents unauthorized access to, or modification of, computer programs or data.

Non-repudiation 
: Degree to which actions or events can be proven to have taken place so that the events or actions cannot be repudiated later.

Accountability 
: Degree to which the actions of an entity can be traced uniquely to the entity.

Authenticity 
: Degree to which the identity of a subject or resource can be proved to be the one claimed.

### Maintainability

This characteristic represents the degree of effectiveness and efficiency with which a product or system can be modified to improve it, correct it or adapt it to changes in environment, and in requirements. This characteristic is composed of the following sub-characteristics:

Modularity 
: Degree to which a system or computer program is composed of discrete components such that a change to one component has minimal impact on other components.

Reusability 
: Degree to which an asset can be used in more than one system, or in building other assets.

Analysability 
: Degree of effectiveness and efficiency with which it is possible to assess the impact on a product or system of an intended change to one or more of its parts, or to diagnose a product for deficiencies or causes of failures, or to identify parts to be modified.

Modifiability
: Degree to which a product or system can be effectively and efficiently modified without introducing defects or degrading existing product quality.

Testability
: Degree of effectiveness and efficiency with which test criteria can be established for a system, product or component and tests can be performed to determine whether those criteria have been met.

### Portability

Degree of effectiveness and efficiency with which a system, product or component can be transferred from one hardware, software or other operational or usage environment to another. This characteristic is composed of the following sub-characteristics:

Adaptability
: Degree to which a product or system can effectively and efficiently be adapted for different or evolving hardware, software or other operational or usage environments.

Installability
: Degree of effectiveness and efficiency with which a product or system can be successfully installed and/or uninstalled in a specified environment.

Replaceability
: Degree to which a product can replace another specified software product for the same purpose in the same environment.
-->


## Quality Scenarios
<!--

**Contents.**

Concretization of (sometimes vague or implicit) quality requirements
using (quality) scenarios.

These scenarios describe what should happen when a stimulus arrives at
the system.

For architects, two kinds of scenarios are important:

  - Usage scenarios (also called application scenarios or use case
    scenarios) describe the system’s runtime reaction to a certain
    stimulus. This also includes scenarios that describe the system’s
    efficiency or performance. Example: The system reacts to a user’s
    request within one second.

  - Change scenarios describe a modification of the system or of its
    immediate environment. Example: Additional functionality is
    implemented or requirements for a quality attribute change.

**Motivation.**

Scenarios make quality requirements concrete and allow to more easily
measure or decide whether they are fulfilled.

Especially when you want to assess your architecture using methods like
ATAM you need to describe your quality goals (from section 1.2) more
precisely down to a level of scenarios that can be discussed and
evaluated.

**Form.**

Tabular or free form text.
-->

