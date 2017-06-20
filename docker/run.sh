#!/bin/bash

#------------------------------------------------------------------------------
# Start HOR API Service :
#------------------------------------------------------------------------------
start_oracle_service_broker() {
    echo 'Starting Oracle-Service-Broker API Service ...'
    bee run -gendoc=true -downdoc=false
}

#------------------------------------------------------------------------------
# HOR API Service Main:
#------------------------------------------------------------------------------
start_oracle_service_broker
