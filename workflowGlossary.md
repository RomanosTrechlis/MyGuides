## Workflow Glossary

**Actor:** user or system that performs actions

**Assignment:** the work a user or a system is ordered to perform

**Case:** any distinct action or process that can be executed

**Conditions:** requirements to be met in order an activity can take place

**Process:** a number of tasks that need to be carried out and a set of conditions that determine the order of the tasks

**Task:** 

1. a logical unit of work that is carried out as a single whole by one resource
2. a process that cannot be subdivided any further: an atomic process. 


[Term Glossary](http://www.wfmc.org/standards/docs/TC-1011_term_glossary_v3.pdf)

**Business Process:** what is intended to happen

**Manual Activities:** are not managed as part of Workflow System

**Process Definition:** a representation of what is intended to happen

**Process Instances:** a representation of what is actually happening

**Work Items:** tasks allocated to a workflow participant

**Workflow:** is the automation of a business process, in whole or part, during which documents, information or tasks are passed from one participant to another for action, according to a set of procedural rules

**Workflow Management System:** controls automated aspects of the business process


### Simplified Relations

A *Business Process* **is** defined **in a** *Process Definition* and it is managed by a *Workflow Management System*.

A *Process Definition* is composed **of** *Activities* which may be *Manual* or *Automated*.

A *Process Definition* is used to create and manage *Process Instances* which include one or more *Activity Instances*.

An *Automated Activity* during execution is represented by *Activity Instances*.

An *Activity Instance* includes *Work Items* and/or *Invoked Applications*.