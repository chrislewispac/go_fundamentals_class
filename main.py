"""
def sortOrders(orderList):
    primes = []
    non_primes = []
    
    # build list of sorted prime orders. build list of non-prime orders in original order
    for i, o in enumerate(orderList):
        order = o.split()
        
        if len(order) <= 1: #handle error of no metadata
            raise Exception('Expected metadata for every order but did not receive metadata for order id', order[0])
        
        order_id, metadata = order[0], order[1:len(order)]
        
        if str.isdigit(metadata[0]):
            non_primes.append([order_id, " ".join(metadata)])
        else:
            primes.append([order_id, " ".join(metadata)])
                   
    # sorted by meta data (alphabetically) 1st and identifier 2nd (ascii value)     
    primes.sort(key = lambda x: (x[1], x[0]))
    
    output = []
    new_list = primes + non_primes
    for item in new_list:
        output.append(" ".join(item))
    
    return  output
"""


"""
import math

def routePairs(maxTravelDist, forwardRouteList, returnRouteList):
    output = [[]]
    curr_best = -math.inf

    #sort both lists by travel dist
    forwardRouteList.sort(key=lambda x: x[1])
    returnRouteList.sort(key=lambda x: x[1])
    
    for i in range(len(forwardRouteList)):
        for j in range(len(returnRouteList)):
            curr_dist = forwardRouteList[i][1] + returnRouteList[j][1]
            if curr_dist <= maxTravelDist:
                #candidate solution
                # if it better than existing solution? if so replace (flush result and add this one)
                if curr_dist > curr_best:
                    curr_best = curr_dist
                    output = [[forwardRouteList[i][0], returnRouteList[j][0]]]
                elif curr_dist == curr_best: #if same as curr_best then add to ouput
                    output.append([forwardRouteList[i][0], returnRouteList[j][0]])
                
    return output
"""