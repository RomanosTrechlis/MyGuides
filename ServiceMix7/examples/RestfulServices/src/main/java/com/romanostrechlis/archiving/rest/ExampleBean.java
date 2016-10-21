/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.romanostrechlis.archiving.rest;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.QueryParam;

/**
 *
 * @author Romanos
 */
@Path("/example")
public class ExampleBean {

    @GET
    @Path("/")
    public String ping() throws Exception {
        return "PONG";
    }
    
    @GET
    @Path("/hello")
    public String greetings(@QueryParam("name") String name) throws Exception {
        return "Hello " + name;
    }
}
