/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.romanostrechlis.incidentreport.service;

import javax.jws.WebService;

/**
 *
 * @author Romanos
 */
@WebService(serviceName = "PingService", endpointInterface = "com.romanostrechlis.incidentreport.service.Ping")
public class PingImpl implements Ping {

    public String ping() {
        return "Pong";
    }
    
}
