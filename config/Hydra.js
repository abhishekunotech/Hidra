{

    "version": "1.0",

    "routes": {

        "Index": {

            "version": "1.0",

            "method": "GET",

            "URI": "/",

            "component": "otrs",

            "handler": "handlers.Index"

        }

        ,

        "GetTicketDetails": {

            "version": "1.0",

            "method": "GET",

            "URI": "/getTicketDetails",

            "component": "otrs",

            "handler": "handlers.GetTicketDetails"

        }

        ,

        "getCILogs": {

            "version": "1.0",

            "method": "GET",

            "URI": "/getCILogs",

            "component": "otrs",

            "handler": "handlers.GetCILogs"

        }

        ,

        "getCIJobs": {

            "version": "1.0",

            "method": "GET",

            "URI": "/getCIJobs",

            "component": "otrs",

            "handler": "handlers.GetCIJobs"

        }

        ,

        "getCIDetails": {

            "version": "1.0",

            "method": "GET",

            "URI": "/getCIDetails",

            "component": "otrs",

            "handler": "handlers.GetCIDetails"

        }

        ,

        "getListOfWorkOrders": {

            "version": "1.0",

            "method": "GET",

            "URI": "/getListOfWorkOrders",

            "component": "otrs",

            "handler": "handlers.GetListOfWorkOrders"

        }

        ,

        "Ticketcreate": {

            "version": "1.0",

            "method": "POST",

            "URI": "/Ticketcreate",

            "component": "otrs",

            "handler": "handlers.Ticketcreate"

        }

        ,

        "getLinkedChange": {

            "version": "1.0",

            "method": "GET",

            "URI": "/getLinkedChange",

            "component": "otrs",

            "handler": "handlers.GetLinkedChange"

        }

        ,

        "getLinkedTickets": {

            "version": "1.0",

            "method": "GET",

            "URI": "/getLinkedTickets",

            "component": "otrs",

            "handler": "handlers.GetLinkedTickets"

        }

        

    }

    ,

    "components": {

        "otrs": {

            "version": "1.0",

            "url": "http://abc.com",

            "apis": {

                "GetTicketDetails": {

                    "version": "1.0",

                    "URI": "/TicketAPI/getTicketDetails",

                    "parameters": {

                        "TicketID": "",

                        "UserLogin": "",

                        "Password": ""

                    }

                    

                }

                

            }

            

        }

        ,

        "glpi": {

            "version": "1.0",

            "url": "",

            "apis": {

                "GetCIDetails": {

                    "version": "1.0",

                    "URI": "/getCIDetails",

                    "parameters": {

                        "TicketID": "",

                        "UserLogin": "",

                        "Password": ""

                    }

                    

                }

                

            }

            

        }

        

    }

    

}


