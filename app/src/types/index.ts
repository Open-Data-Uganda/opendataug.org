export type Customer = {
  generated_customer_number: string;
  number: string;
  name: string;
  postal_code: string;
  city: string;
  phone_number: string;
  email_address: string;
  status: string;
};

export type User = {
  number: string;
  name: string;
  email: string;
  first_name: string;
  other_name: string;
  phone: string;
  role: string;
  status: string;
};

export type Quotation = {
  created_at: string;
  updated_at: string;
  quotation_total: number;
  generated_quotation_number: string;
  number: string;
  customer: string;
  building_type: string;
  status: string;
  doors_needed: string;
  no_of_bays: number;
  bay_size: number;
  building_width: number;
  building_width_units: string;
  building_height: number;
  building_height_units: string;
  roof_shape: string;
  bay_span: number;
  bay_length: number;
  building_span: number;
  total_bays_in_span: number;
  total_bays_in_length: number;
  bay_quantity: number;
  eaves: number;
  purlin_spacing: number;
  purlin_length: number;
  purlin_quantity: number;
  column_height: number;
  column_quantity: number;
  column_girth: number;
  column_type: string;
  roof_mono_pitch_degrees: number;
  roof_pitch_standard: number;
  rafter: number;
  rafter_quantity: number;
  gable_post: number;
  rafter_stays_quantity: number;
  rafter_stays_positions: string;
  foundation_quantity: number;
  guttering: number;
  eaves_beam: number;
  wind_bracing_roof_length: number;
  wind_bracing_roof_quantity: number;
  wind_bracing_sides_length: number;
  wind_bracing_sides_quantity: number;
  wind_bracing_total_length: number;
  height_to_ridge_on_span: number;
  height_to_eaves_on_length: number;
  base_plate_dimensions: string;
  base_plate_quantity: number;
  concrete_dimensions: number;
  column_restraints_quantity: number;
  rafter_restraints_quantity: number;
  cladding_rail_length: number;
  end_plate_quantity: number;
  cleat_plate_quantity: number;
  stiffeners_quantity: number;
  angles_quantity: number;
  channels_quantity: number;
  sleeves_quantity: number;
  tie_wires_quantity: number;
  bay_side_rails_length: number;
  cleader_rails_length: number;
  facia_boards_quantity: number;
  flashings_quantity: number;
  gates: number;
  door_quantity: number;
  door_type: string;
  door_width: number;
  door_tracks: number;
  door_wheels: number;
  feed_barrier_tops_length: number;
  feed_barrier_timber_base_length: number;
  concrete_joining_corners: number;
};

export type MaterialPrice = {
  number: string;
  unit_cost: number;
  name: string;
  material_type: string;
  quantity: number;
  price: number;
  description: string;
  status: string;
  unit_of_measure: string;
  length: number;
  width: number;
  height: number;
  dimensions_unit: string;
};

export type Invoice = {
  generated_invoice_number: string;
  number: string;
  customer: string;
  quotation: string;
  total: number;
  status: string;
  due_date: string;
  issue_date: string;
};

export const invoiceStatuses = [
  { value: '', label: 'All' },
  { value: 'PENDING', label: 'Pending' },
  { value: 'PAID', label: 'Paid' },
  { value: 'OVERDUE', label: 'Overdue' }
];

export const invoiceStatusesMinor = [
  { value: 'PENDING', label: 'Pending' },
  { value: 'PAID', label: 'Paid' },
  { value: 'OVERDUE', label: 'Overdue' }
];

export const materialFilterStatuses = [
  { value: '', label: 'All' },
  { value: 'CHECK_AVAILABILITY', label: 'Check Availability' },
  { value: 'REQUEST_SUPPLIER_QUOTE', label: 'Request Supplier Quote' },
  { value: 'PENDING_DELIVERY', label: 'Pending Delivery' },
  { value: 'UNAVAILABLE', label: 'Unavailable' },
  { value: 'REQUIRES_ATTENTION', label: 'Requires Attention' },
  { value: 'CHECK_TIMBER_STORAGE', label: 'Check Timber Storage' },
  { value: 'CHECK_FIXTURES_FITTINGS', label: 'Check Fixtures + Fittings' },
  { value: 'BESPOKE_PARTS_REQUIRED', label: 'Bespoke Parts Required' }
];

export const materialStatuses = [
  { value: 'CHECK_AVAILABILITY', label: 'Check Availability' },
  { value: 'REQUEST_SUPPLIER_QUOTE', label: 'Request Supplier Quote' },
  { value: 'PENDING_DELIVERY', label: 'Pending Delivery' },
  { value: 'UNAVAILABLE', label: 'Unavailable' },
  { value: 'REQUIRES_ATTENTION', label: 'Requires Attention' },
  { value: 'CHECK_TIMBER_STORAGE', label: 'Check Timber Storage' },
  { value: 'CHECK_FIXTURES_FITTINGS', label: 'Check Fixtures + Fittings' },
  { value: 'BESPOKE_PARTS_REQUIRED', label: 'Bespoke Parts Required' }
];

export const quotationFilterStatuses = [
  { value: '', label: 'All' },
  { value: 'DRAFT', label: 'Draft' },
  { value: 'PENDING_REVIEW', label: 'Pending Review' },
  { value: 'OUTSTANDING', label: 'Outstanding' },
  { value: 'PENDING_APPROVAL', label: 'Pending Approval' },
  { value: 'DEPOSIT_REQUIRED', label: 'Deposit Required' },
  { value: 'DEPOSIT_PAID', label: 'Deposit Paid' },
  { value: 'AWAITING_PARTS', label: 'Awaiting Parts' },
  { value: 'AWAITING_DELIVERY', label: 'Awaiting Delivery' },
  { value: 'ARRANGE_DELIVERY', label: 'Arrange Delivery' },
  { value: 'ARRANGE_INVOICE', label: 'Arrange Invoice' },
  { value: 'PENDING_PAYMENT', label: 'Pending Payment' },
  { value: 'ARCHIVED', label: 'Archived' }
];

export const quotationStatuses = [
  { value: 'DRAFT', label: 'Draft' },
  { value: 'PENDING_REVIEW', label: 'Pending Review' },
  { value: 'OUTSTANDING', label: 'Outstanding' },
  { value: 'PENDING_APPROVAL', label: 'Pending Approval' },
  { value: 'DEPOSIT_REQUIRED', label: 'Deposit Required' },
  { value: 'DEPOSIT_PAID', label: 'Deposit Paid' },
  { value: 'AWAITING_PARTS', label: 'Awaiting Parts' },
  { value: 'AWAITING_DELIVERY', label: 'Awaiting Delivery' },
  { value: 'ARRANGE_DELIVERY', label: 'Arrange Delivery' },
  { value: 'ARRANGE_INVOICE', label: 'Arrange Invoice' },
  { value: 'PENDING_PAYMENT', label: 'Pending Payment' },
  { value: 'ARCHIVED', label: 'Archived' }
];

export const customerStatuses = [
  { value: '', label: 'All' },
  { value: 'ACTIVE', label: 'Active' },
  { value: 'INACTIVE', label: 'Inactive' }
];

export const editCustomerStatuses = [
  { value: 'ACTIVE', label: 'Active' },
  { value: 'INACTIVE', label: 'Inactive' }
];

export const buildingTypes = [
  'Cattle Building',
  'Straw Shed',
  'Grain Store',
  'Dung Store',
  'Storage Shed',
  'Sheep Shed',
  'Workshop',
  'Stable Building',
  'Warehouse',
  'Custom'
];

export const meseauringUnits = ['Feet', 'Metres'];

export const generalMeasuringUnits = ['Inches', 'Millimeters', 'Feet', 'Metres'];

export const smallerMeseauringUnits = ['Inches', 'Millimeters'];

export const smallerMeseauringUnitsOptional = ['Feet', 'Millimeters'];

export const sizeOfBays = ['4.572 m', '6.1 m', 'Other'];

export const steelTypes = ['Hot Dipped Galvanised', 'Primer Painted', 'Gloss Finish Painted', 'No Finish'];

export const paintColours = ['Red', 'Grey'];

export const roofShapes = ['Apex Roof', 'Mono-pitch Roof'];

export const roofColors = [
  'Anthracite Grey',
  'Light Grey',
  'Juniper Green',
  'Olive Green',
  'Moorland Green',
  'Slate Blue',
  'Vandyke Brown',
  'Black'
];

export const ridgeTypes = ['50% Vented', 'Closed', 'Open Protected'];

export const purlinsDimensions = ['225mm x 75mm', '175mm x 75mm'];

export const doorTypes = [
  'Sliding Doors',
  'Swinging Hinge Barn Door',
  'Roller Shutter Door',
  'Single Personnel Security Door'
];

export const doorLocation = ['Front', 'Back', 'Left Side', 'Right Side'];

export const doorBayLocation = ['First Bay', 'Second Bay', 'Third Bay', 'Forth Bay', 'Fifth Bay', 'Etc'];

export const specialOffers = [
  '24.4 x 15.24 x 4.8 Dung Store',
  '30.48 x 15.24 x 6.4 Straw Shed',
  '30.48 x 12.192 x 4.572 Cattle Shed with 1.37 Cantilever'
];

export const roofColorOffers = [
  'Anthracite Grey',
  'Light Grey',
  'Juniper Green',
  'Olive Green',
  'Moorland Green',
  'Slate Blue',
  'Vandyke Brown',
  'Black'
];

export const claddingColors = [
  'Anthracite Grey',
  'Light Grey',
  'Juniper Green',
  'Olive Green',
  'Moorland Green',
  'Slate Blue',
  'Vandyke Brown',
  'Black'
];

export const claddingLocations = ['Front', 'Back', 'Left Side', 'Right Side'];

export const wallLocation = ['Front', 'Back', 'Left Side', 'Right Side'];

export const gutterringTypes = ['170mm High Flow', '200mm Storm Flow'];

export const cleaderRailType = ['Timber', 'Steel 100x100x2mm'];

export const claddingMaterials = [
  'Fibre Cement',
  'Box Profile Tin',
  'Box Profile Tin with Anti-Condensation',
  'Composite Panel 40mm Thickness',
  'Composite Panel 60mm Thickness',
  'Composite Panel 80mm Thickness',
  'Composite Panel 100mm Thickness',
  'Composite Panel 120mm Thickness',
  'Yorkshire Boarding'
];

export const mainCladdingLocations = [
  'Yorkshire Boarding',
  'Double-Lap Yorkshire Boarding',
  'Box Profile Tin',
  'Composite Panels'
];

export const mainCladdingColors = [
  'Anthracite Grey',
  'Light Grey',
  'Juniper Green',
  'Olive Green',
  'Moorland Green',
  'Slate Blue',
  'Vandyke Brown',
  'Black'
];

export const wallMaterials = ['Pre-Cast Concrete Panels', 'Stock Boarding'];
export const concreteWallThicknessTypes = ['150mm thick concrete panel wall', '100mm thick concrete panel wall'];

export const wallTypes = ['Pre-Cast Concrete Panels', 'Stock Boarding'];

export const sourceOfInfo = ['Social Media', 'Internet Search', 'Friend', 'Advertisement', 'Other'];

export const projectTime = ['Weeks', 'Months'];

export const buildingUse = [
  'Cattle Building',
  'Straw Shed',
  'Grain Store',
  'Dung Store',
  'Storage Shed',
  'Sheep Shed',
  'Workshop',
  'Stable Building',
  'Warehouse',
  'Custom'
];

export const finishingDetails = ['Interior Concrete', 'Exterior Concrete'];

export const steps = [
  {
    id: 'Step 1',
    name: 'Personal Information',
    fields: ['first_name', 'other_name', 'email_address', 'phone_number']
  },
  {
    id: 'Step 2',
    name: 'Address details',
    fields: ['street_address', 'alternative_street_address', 'city', 'province', 'postal_code']
  },
  {
    id: 'Step 3',
    name: 'Structure Information',
    fields: [
      'building_type',
      'building_height',
      'building_height_units',
      'building_width',
      'building_width_units',
      'no_of_bays',
      'bay_size',
      'roof_shape',
      'roof_cladding_material',
      'roof_color',
      'ridge_type',
      'guttering_needed',
      'guttering_type'
    ]
  },
  {
    id: 'Step 4',
    name: 'Elevation One',
    fields: [
      'wall_one_walls_needed',
      'wall_one_cladding_needed',
      'wall_one_height',
      'wall_one_length',
      'wall_one_length_units',
      'wall_one_height_units',
      'wall_one_thickness',
      'wall_one_thickness_units',
      'wall_one_material',
      'wall_one_cladding_type',
      'wall_one_cladding_colour',
      'wall_one_door_quantity',
      'wall_one_doors_needed',
      'wall_one_door_type',
      'wall_one_door_height',
      'wall_one_door_height_units',
      'wall_one_door_width',
      'wall_one_door_width_units',
      'wall_one_cantilever_needed',
      'wall_one_cantilever'
    ]
  },

  {
    id: 'Step 5',
    name: 'Elevation Two',
    fields: [
      'wall_two_walls_needed',
      'wall_two_cladding_needed',
      'wall_two_height',
      'wall_two_length',
      'wall_two_length_units',
      'wall_two_height_units',
      'wall_two_thickness',
      'wall_two_thickness_units',
      'wall_two_material',
      'wall_two_cladding_type',
      'wall_two_cladding_colour',
      'wall_two_door_quantity',
      'wall_two_doors_needed',
      'wall_two_door_type',
      'wall_two_door_height',
      'wall_two_door_height_units',
      'wall_two_door_width',
      'wall_two_door_width_units',
      'wall_two_cantilever_needed',
      'wall_two_cantilever'
    ]
  },
  {
    id: 'Step 6',
    name: 'Elevation Three',
    fields: [
      'wall_three_walls_needed',
      'wall_three_cladding_needed',
      'wall_three_height',
      'wall_three_length',
      'wall_three_length_units',
      'wall_three_height_units',
      'wall_three_thickness',
      'wall_three_thickness_units',
      'wall_three_material',
      'wall_three_cladding_type',
      'wall_three_cladding_colour',
      'wall_three_door_quantity',
      'wall_three_doors_needed',
      'wall_three_door_type',
      'wall_three_door_height',
      'wall_three_door_height_units',
      'wall_three_door_width',
      'wall_three_door_width_units',
      'wall_three_cantilever_needed',
      'wall_three_cantilever'
    ]
  },
  {
    id: 'Step 7',
    name: 'Elevation Four',
    fields: [
      'wall_four_walls_needed',
      'wall_four_cladding_needed',
      'wall_four_height',
      'wall_four_length',
      'wall_four_length_units',
      'wall_four_height_units',
      'wall_four_thickness',
      'wall_four_thickness_units',
      'wall_four_material',
      'wall_four_cladding_type',
      'wall_four_cladding_colour',
      'wall_four_door_quantity',
      'wall_four_doors_needed',
      'wall_four_door_type',
      'wall_four_door_height',
      'wall_four_door_height_units',
      'wall_four_door_width',
      'wall_four_door_width_units',
      'wall_four_cantilever_needed',
      'wall_four_cantilever'
    ]
  },

  { id: 'Step 8', name: 'Complete' }
];

export const quotationDetailsColumns = [
  {
    id: 1,
    name: 'Material'
  },
  {
    id: 2,
    name: 'Type'
  },
  {
    id: 3,
    name: 'Sub-Type'
  },
  {
    id: 4,
    name: 'Unit of Measurement'
  },
  {
    id: 5,
    name: 'Quantity'
  },
  {
    id: 6,
    name: 'Price per Unit'
  },
  {
    id: 7,
    name: 'Total'
  }
];

export type MaterialType = {
  name: string;
  description: string;
  Quantity: number;
  status: string;
  manufacturer: string;
  lead_time: string;
  unit_cost: number;
  weight: number;
  colour: string;
  material_model: string;
  image_url: string;
  material_type: string;
  unit_of_measure: string;
  number: string;
};

export const roleTypes = [
  { value: 'STAFF', label: 'Staff' },
  { value: 'ADMIN', label: 'Administrator' }
];

export const userStatuses = [
  { value: 'ACTIVE', label: 'Active' },
  { value: 'INACTIVE', label: 'Inactive' },
  { value: 'DEACTIVATE', label: 'Deactivate' }
];

export const gablePostDimensionTypes = [
  { value: '152x89x16', label: '152 x 89 x 16' },
  { value: '203x102x23', label: '203 x 102 x 23' },
  { value: '203x133x30', label: '203 x 133 x 30' },
  { value: '127x76x13', label: '127 x 76 x 13' },
  { value: '203x133x25', label: '203 x 133 x 25' },
  { value: '178x102x19', label: '178 x 102 x 19' },
  { value: '254x146x37', label: '254 x 146 x 37' },
  { value: '254x146x31', label: '254 x 146 x 31' },
  { value: '305x165x40', label: '305 x 165 x 40' },
  { value: '305x165x46', label: '305 x 165 x 46' },
  { value: '254x146x43', label: '254 x 146 x 43' }
];

export const foundationSizes = [
  { value: 'A-1x1x0.5', label: 'A-1 x 1 x 0.5' },
  { value: 'B-1x1x0.6', label: 'B-1 x 1 x 0.6' },
  { value: 'C-1.2x1.2x0.6', label: 'C-1.2 x 1.2 x 0.6' }
];

export const columnSizes = [
  { value: '178x102x19', label: '178 x 102 x 19' },
  { value: '203x133x25', label: '203 x 133 x 25' },
  { value: '254x146x31', label: '254 x 146 x 31' },
  { value: '356x127x33', label: '356 x 127 x 33' },
  { value: '356x127x39', label: '356 x 127 x 39' },
  { value: '356x171x45', label: '356 x 171 x 45' },
  { value: '254x146x37', label: '254 x 146 x 37' },
  { value: '305x165x40', label: '305 x 165 x 40' },
  { value: '305x165x46', label: '305 x 165 x 46' },
  { value: '356x171x51', label: '356 x 171 x 51' },
  { value: '406x178x54', label: '406 x 178 x 54' },
  { value: '406x178x60', label: '406 x 178 x 60' },
  { value: '406x178x74', label: '406 x 178 x 74' }
];

export const rafterSizes = [
  { value: '152x89x16', label: '152 x 89 x 16' },
  { value: '178x102x19', label: '178 x 102 x 19' },
  { value: '203x102x23', label: '203 x 102 x 23' },
  { value: '203x133x25', label: '203 x 133 x 25' },
  { value: '203x133x30', label: '203 x 133 x 30' },
  { value: '254x146x31', label: '254 x 146 x 31' },
  { value: '254x146x37', label: '254 x 146 x 37' },
  { value: '305x165x40', label: '305 x 165 x 40' },
  { value: '356x127x33', label: '356 x 127 x 33' },
  { value: '356x127x39', label: '356 x 127 x 39' },
  { value: '406x140x39', label: '406 x 140 x 39' },
  { value: '406x140x46', label: '406 x 140 x 46' },
  { value: '356x171x45', label: '356 x 171 x 45' }
];

export const purlinTypes = [
  { value: 'Timber', label: 'Timber' },
  { value: 'Steel', label: 'Steel' }
];

export const windBracingTypes = [
  { value: 12, label: 12 },
  { value: 15, label: 15 }
];

export const eavesBeam = [
  { value: '225x75mm Eaves Beam - Timber', label: '225x75mm Eaves Beam - Timber' },
  { value: '200mm x 90mm Flat Faced Eaves Beam', label: '200mm x 90mm Flat Faced Eaves Beam' }
];

export const dimensionTypes = [
  { value: '178x102x19', label: '178 x 102 x 19' },
  { value: '203x133x25', label: '203 x 133 x 25' },
  { value: '254x146x31', label: '254 x 146 x 31' },
  { value: '203x133x30', label: '203 x 133 x 30' },
  { value: '356x127x33', label: '356 x 127 x 33' },
  { value: '305x165x40', label: '305 x 165 x 40' },
  { value: '356x127x39', label: '356 x 127 x 39' },
  { value: '356x171x45', label: '356 x 171 x 45' },
  { value: '254x146x37', label: '254 x 146 x 37' },
  { value: '305x165x46', label: '305 x 165 x 46' },
  { value: '406x178x54', label: '406 x 178 x 54' },
  { value: '356x171x51', label: '356 x 171 x 51' },
  { value: '406x178x60', label: '406 x 178 x 60' },
  { value: '406x178x74', label: '406 x 178 x 74' }
];

export const basePlateType = ['300 x 300mm', '400 x 300mm'];

export const sideRailMaterials = ['Timber', 'Cold-rolled Steel'];

export const sideRailTimberTypes = ['125x75mm', '150x75mm'];
export const sideRailColdRolledSteelTypes = ['175x65x1.6mm', '200x65x1.6mm'];

export const roofRidgeOptions = {
  boxProfile: ['Vented', 'Sealed'],
  compositePanel: ['Sealed']
};
