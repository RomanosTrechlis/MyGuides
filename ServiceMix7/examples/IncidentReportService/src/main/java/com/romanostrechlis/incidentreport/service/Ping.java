/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.romanostrechlis.incidentreport.service;

import javax.jws.WebMethod;
import javax.jws.WebService;
import javax.xml.bind.annotation.XmlSeeAlso;

/**
 *
 * @author Romanos
 */
@WebService
@XmlSeeAlso({com.romanostrechlis.incidentreport.service.types.ObjectFactory.class})
public interface Ping {
    
    @WebMethod(operationName = "Ping")
    public String ping();
}
