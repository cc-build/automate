export interface DataBag {
  server_id: string;
  org_id: string;
  name: string;
}

export interface DataBagItems {
  name: string;
  active?: boolean;
}

export interface DataBagsItemDetails {
  server_id: string;
  org_id: string;
  data_bag_name: string;
  name: string;
  id: string;
  data: string;
}
